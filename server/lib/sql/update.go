package sqli

import (
	lib "Spark/server/lib/func"
	"database/sql"
)

func UpdateQuery(db *sql.DB, table string, prefix string, upCol []string, upVal []string, condCol []string, condVal []string,
	pad string, extraString string) (int64, error) {
	// 拼接更新语句
	var sqlQuery = "UPDATE `" + prefix + table + "` SET " +
		sqlUpdateConnect(upCol) +
		" WHERE " + sqlConditionConnect(condCol, pad) + extraString
	// 使用预处理，防止SQL注入
	stmt, err := db.Prepare(sqlQuery)
	if err != nil {
		return 0, err
	}
	// 获取预处理得到的结果集
	result, err := stmt.Exec(lib.MergeSliceStringValue(upVal, condVal)...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
