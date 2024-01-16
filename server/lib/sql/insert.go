package sqli

import (
	lib "Spark/server/lib/func"
	"database/sql"
)

func InsertSafe(db *sql.DB, table string, prefix string, insCol []string, insVal []string, extraString string) (int64, error) {
	// 拼接插入语句
	strCol, strVal := sqlInsertConnect(insCol)
	var sqlQuery = "INSERT INTO `" + prefix + table + "`" + strCol + " VALUES" + strVal + extraString
	// 使用预处理，防止SQL注入
	stmt, err := db.Prepare(sqlQuery)
	if err != nil {
		return 0, err
	}
	// 获取预处理得到的结果集
	result, err := stmt.Exec(lib.MergeSliceStringValue(insVal)...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
