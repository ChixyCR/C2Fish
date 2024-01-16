package lib

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"golang.org/x/crypto/sha3"
	"hash"
	"strings"
)

func GetHashSignature(s ,e []byte,t string,x bool) interface{}{

	var h interface{} = nil
	// 判断需要进行哈希的算法
	switch strings.ToLower(t) {
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "sha224":
		h = sha3.New224()
	case "sha256":
		h = sha256.New()
	case "sha384":
		h = sha512.New384()
	case "sha512":
		h = sha512.New()
	default:
		return nil
	}
	// 进行哈希运算
	h.(hash.Hash).Write(append(s,e...))
	// 输出2进制切片或hex字符
	if x {
		return hex.EncodeToString(h.(hash.Hash).Sum(nil))
	}
	return h.(hash.Hash).Sum(nil)
}