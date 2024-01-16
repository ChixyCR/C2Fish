package bridge

import (
	"Spark/modules"
	"Spark/utils"
	"Spark/utils/cmap"
	"github.com/gin-gonic/gin"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

// Bridge is a utility that handles the binary flow from the client
// to the browser or flow from the browser to the client.

type Bridge struct {
	creation int64
	using    bool
	uuid     string
	lock     *sync.Mutex
	Dst      *gin.Context
	Src      *gin.Context
	ext      any
	OnPull   func(bridge *Bridge)
	OnPush   func(bridge *Bridge)
	OnFinish func(bridge *Bridge)
}

var bridges = cmap.New[*Bridge]()

func init() {
	go func() {
		for now := range time.NewTicker(15 * time.Second).C {
			var queue []string
			timestamp := now.Unix()
			bridges.IterCb(func(k string, b *Bridge) bool {
				if timestamp-b.creation > 60 && !b.using {
					b.lock.Lock()
					if b.Src != nil && b.Src.Request.Body != nil {
						b.Src.Request.Body.Close()
					}
					b.Src = nil
					b.Dst = nil
					b.lock.Unlock()
					b = nil
					queue = append(queue, b.uuid)
				}
				return true
			})
			bridges.Remove(queue...)
		}
	}()
}

func CheckBridge(ctx *gin.Context) *Bridge {
	var form struct {
		Bridge string `json:"bridge" yaml:"bridge" form:"bridge" binding:"required"`
	}
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, modules.Packet{Code: -1, Msg: `${i18n|COMMON.INVALID_PARAMETER}`})
		return nil
	}
	b, ok := bridges.Get(form.Bridge)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, modules.Packet{Code: -1, Msg: `${i18n|COMMON.INVALID_BRIDGE_ID}`})
		return nil
	}
	return b
}

func BridgePush(ctx *gin.Context) {
	bridge := CheckBridge(ctx)
	if bridge == nil {
		return
	}
	bridge.lock.Lock()
	if bridge.using || (bridge.Src != nil && bridge.Dst != nil) {
		bridge.lock.Unlock()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, modules.Packet{Code: 1, Msg: `${i18n|COMMON.BRIDGE_IN_USE}`})
		return
	}
	bridge.Src = ctx
	bridge.using = true
	bridge.lock.Unlock()
	if bridge.OnPush != nil {
		bridge.OnPush(bridge)
	}
	if bridge.Dst != nil && bridge.Dst.Writer != nil {
		// Get net.Conn to set deadline manually.
		SrcConn, SrcOK := bridge.Src.Request.Context().Value(`Conn`).(net.Conn)
		DstConn, DstOK := bridge.Dst.Request.Context().Value(`Conn`).(net.Conn)
		if SrcOK && DstOK {
			for {
				eof := false
				buf := make([]byte, 2<<14)
				SrcConn.SetReadDeadline(utils.Now.Add(5 * time.Second))
				n, err := bridge.Src.Request.Body.Read(buf)
				if n == 0 {
					break
				}
				if err != nil {
					eof = err == io.EOF
					if !eof {
						break
					}
				}
				DstConn.SetWriteDeadline(utils.Now.Add(10 * time.Second))
				_, err = bridge.Dst.Writer.Write(buf[:n])
				if eof || err != nil {
					break
				}
			}
		}
		SrcConn.SetReadDeadline(time.Time{})
		DstConn.SetWriteDeadline(time.Time{})
		bridge.Src.Status(http.StatusOK)
		if bridge.OnFinish != nil {
			bridge.OnFinish(bridge)
		}
		RemoveBridge(bridge.uuid)
		bridge = nil
	}
}

func BridgePull(ctx *gin.Context) {
	bridge := CheckBridge(ctx)
	if bridge == nil {
		return
	}
	bridge.lock.Lock()
	if bridge.using || (bridge.Src != nil && bridge.Dst != nil) {
		bridge.lock.Unlock()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, modules.Packet{Code: 1, Msg: `${i18n|COMMON.BRIDGE_IN_USE}`})
		return
	}
	bridge.Dst = ctx
	bridge.using = true
	bridge.lock.Unlock()
	if bridge.OnPull != nil {
		bridge.OnPull(bridge)
	}
	if bridge.Src != nil && bridge.Src.Request.Body != nil {
		// Get net.Conn to set deadline manually.
		SrcConn, SrcOK := bridge.Src.Request.Context().Value(`Conn`).(net.Conn)
		DstConn, DstOK := bridge.Dst.Request.Context().Value(`Conn`).(net.Conn)
		if SrcOK && DstOK {
			for {
				eof := false
				buf := make([]byte, 2<<14)
				SrcConn.SetReadDeadline(utils.Now.Add(5 * time.Second))
				n, err := bridge.Src.Request.Body.Read(buf)
				if n == 0 {
					break
				}
				if err != nil {
					eof = err == io.EOF
					if !eof {
						break
					}
				}
				DstConn.SetWriteDeadline(utils.Now.Add(10 * time.Second))
				_, err = bridge.Dst.Writer.Write(buf[:n])
				if eof || err != nil {
					break
				}
			}
		}
		SrcConn.SetReadDeadline(time.Time{})
		DstConn.SetWriteDeadline(time.Time{})
		bridge.Src.Status(http.StatusOK)
		if bridge.OnFinish != nil {
			bridge.OnFinish(bridge)
		}
		RemoveBridge(bridge.uuid)
		bridge = nil
	}
}

func AddBridge(ext any, uuid string) *Bridge {
	bridge := &Bridge{
		creation: utils.Unix,
		uuid:     uuid,
		using:    false,
		lock:     &sync.Mutex{},
		ext:      ext,
	}
	bridges.Set(uuid, bridge)
	return bridge
}

func AddBridgeWithSrc(ext any, uuid string, Src *gin.Context) *Bridge {
	bridge := &Bridge{
		creation: utils.Unix,
		uuid:     uuid,
		using:    false,
		lock:     &sync.Mutex{},
		ext:      ext,
		Src:      Src,
	}
	bridges.Set(uuid, bridge)
	return bridge
}

func AddBridgeWithDst(ext any, uuid string, Dst *gin.Context) *Bridge {
	bridge := &Bridge{
		creation: utils.Unix,
		uuid:     uuid,
		using:    false,
		lock:     &sync.Mutex{},
		ext:      ext,
		Dst:      Dst,
	}
	bridges.Set(uuid, bridge)
	return bridge
}

func RemoveBridge(uuid string) {
	b, ok := bridges.Get(uuid)
	if !ok {
		return
	}
	bridges.Remove(uuid)
	if b.Src != nil && b.Src.Request.Body != nil {
		b.Src.Request.Body.Close()
	}
	b.Src = nil
	b.Dst = nil
	b = nil
}
