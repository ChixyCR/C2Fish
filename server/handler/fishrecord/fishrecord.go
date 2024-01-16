package fishrecord

import (
	sqli "Spark/server/lib/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Fishrecord(ctx *gin.Context) {
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()
	taskID := ctx.Query("taskID") // 获取传入的recordID参数

	col := []string{"recordID", "taskID", "moduleID", "modulePublic", "getMethod", "getIP", "getResult", "getModuleResult", "getTime"}
	condCol := []string{"taskID"} // 查询条件字段为recordID
	condVal := []string{taskID}   // 查询条件值为传入的recordID参数
	pad := ""                     // 查询的limit条件，我这里为空，你可以根据需求自定义
	extraString := ""             // 额外的查询条件，我这里为空，你可以根据需求自定义

	result, err := thisSql.Query("appTaskView", col, condCol, condVal, pad, extraString)
	if err != nil {
		// 处理错误...
		return
	}

	// 将查询结果转为JSON格式返回
	ctx.JSON(http.StatusOK, result)
}
