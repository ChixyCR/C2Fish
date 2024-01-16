package fishstop

import (
	sqli "Spark/server/lib/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Fishstop(ctx *gin.Context) {
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()

	// 获取传入的 taskID 参数
	taskID := ctx.Query("taskID")
	if taskID == "" {
		// 处理错误，taskID 不能为空
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "taskID cannot be empty"})
		return
	}

	// 定义查询字段和条件
	col := []string{"taskID", "taskStatus"}
	condCol := []string{"taskID"} // 查询条件字段为taskID
	condVal := []string{taskID}   // 查询条件值为传入的taskID参数
	pad := ""
	extraString := ""

	// 调用 thisSql.Query 进行查询
	result, err := thisSql.Query("appTasks", col, condCol, condVal, pad, extraString)
	if err != nil {
		// 处理错误...
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 检查是否找到匹配的记录
	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No records found for the given taskID"})
		return
	}

	// 获取当前的 taskStatus 值
	taskStatusValue, exists := result[0]["taskStatus"]
	if !exists {
		// 处理字段不存在的情况
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "taskStatus field not found in the query result"})
		return
	}

	// 将 taskStatusValue 转换为 int64
	currentTaskStatus, ok := taskStatusValue.(int64)
	if !ok {
		// 处理类型断言失败的情况
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assert taskStatus as int64"})
		return
	}

	// 设置更新字段和值
	var newTaskStatus int64
	if currentTaskStatus == 0 {
		newTaskStatus = 1
	} else if currentTaskStatus == 1 {
		newTaskStatus = 0
	} else {
		// 处理无效的 taskStatus 值
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid taskStatus value"})
		return
	}

	updateColumns := []string{"taskStatus"}
	updateValues := []string{strconv.FormatInt(newTaskStatus, 10)}

	// 设置查询条件字段和值
	conditionColumns := []string{"taskID"}
	conditionValues := []string{taskID}

	// 设置其他参数，可以为空
	pad = ""
	extraString = ""

	// 调用 thisSql.Update 进行更新操作
	rowsAffected, err := thisSql.Update("appTasks", updateColumns, updateValues, conditionColumns, conditionValues, pad, extraString)
	if err != nil {
		// 处理错误...
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 检查是否更新成功
	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No records found for the given taskID"})
		return
	}

	// 返回成功的响应
	ctx.JSON(http.StatusOK, gin.H{"message": "Task status updated successfully"})
}
