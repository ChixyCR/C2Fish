package generatefish

import (
	lib "Spark/server/lib/func"
	sqli "Spark/server/lib/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateProject(ctx *gin.Context) {
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()
	taskName := ctx.DefaultPostForm("taskName", "")
	taskData := ctx.DefaultPostForm("taskData", "")
	taskStatus := ctx.DefaultPostForm("taskStatus", "")
	moduleID := ctx.DefaultPostForm("moduleID", "")

	fUrl := ctx.DefaultPostForm("fishUrl", "") //获取TaskName，TaskData，taskStatus，ModuleID,furl

	fileName := ""
	filePath := ""

	// check

	if !lib.CheckTaskName(taskName) {
		ctx.JSON(
			http.StatusOK,
			map[string]string{
				"status":   "error",
				"contents": "The task name is not valid",
			},
		)
		return
	}

	if taskStatus != "1" && taskStatus != "0" {
		taskStatus = "0"
	}

	// Get data

	r, _ := thisSql.Query(
		"appModule",
		[]string{"moduleName"},
		[]string{"moduleID"},
		[]string{moduleID},
		" AND ",
		"",
	)

	if len(r) != 1 {
		ctx.JSON(
			http.StatusOK,
			map[string]string{
				"status":   "error",
				"contents": "Module error",
			},
		)
		return
	}

	// insert

	jsonData, _ := json.Marshal(map[string]interface{}{
		"fishUrl":  fUrl,
		"fishFILE": fileName,
	})

	rt, _ := thisSql.Insert(
		"appTasks",
		[]string{"taskID",
			"taskModuleID",
			"taskName",
			"taskModuleName",
			"taskModulePublic",
			"taskData",
			"taskCode",
			"taskParams",
			"taskStatus",
			"taskRecordNum",
			"filePath",
			"taskCreateTime"},
		[]string{"0",
			moduleID,
			taskName,
			r[0]["moduleName"].(string),
			"1",
			taskData,
			"",
			string(jsonData),
			taskStatus,
			"0",
			filePath,
			lib.GetTimeDateTime(),
		},
		"",
	)

	if rt != 1 {
		ctx.JSON(
			http.StatusOK,
			map[string]string{
				"status":   "error",
				"contents": "New task failed",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		map[string]string{
			"status":   "refresh",
			"contents": "New task successfully created~",
		},
	)
}
