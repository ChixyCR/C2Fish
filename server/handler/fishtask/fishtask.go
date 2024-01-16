package fishtask

import (
	sqli "Spark/server/lib/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Fishtask(ctx *gin.Context) {
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()

	taskID := ctx.Query("taskID") // 获取传入的recordID参数

	col := []string{"taskName", "taskID", "taskData", "taskStatus", "taskRecordNum", "taskParams"}
	condCol := []string{} // 查询条件字段为空切片，表示没有具体的条件
	condVal := []string{} // 查询条件值为空切片，表示没有具体的条件
	pad := ""             // 查询的limit条件，我这里为空，你可以根据需求自定义
	extraString := ""     // 额外的查询条件，我这里为空，你可以根据需求自定义

	// 如果 taskID 存在，添加查询条件
	if taskID != "" {
		condCol = append(condCol, "taskID")
		condVal = append(condVal, taskID)
	}

	result, err := thisSql.Query("appTasks", col, condCol, condVal, pad, extraString)
	if err != nil {
		// 处理错误...
		return
	}

	// 将查询结果转为JSON格式返回
	ctx.JSON(http.StatusOK, result)
}

// Rest of the code remains unchanged
