package sqli

import (
	lib "Spark/server/lib/func"
	"database/sql"
)

func QuerySafe(db *sql.DB, table string, prefix string, col []string, condCol []string, condVal []string, pad string,
	extraString string) (*sql.Rows, []string, error) {
	// 指定字段查询
	if len(condCol) > 0 {
		// 拼接查询语句
		var sqlQuery = "SELECT " +
			sqlColConnect(col) +
			" FROM `" + prefix + table + "` WHERE " +
			sqlConditionConnect(condCol, pad) + extraString
		// 使用参数绑定，防止SQL注入
		rows, err := db.Query(sqlQuery, lib.MergeSliceStringValue(condVal)...)
		if err != nil {
			return nil, nil, err
		}
		// 获取查询结果字段集
		colArr, err := rows.Columns()
		if err != nil {
			return nil, nil, err
		}
		return rows, colArr, err
		// 全字段查询
	} else {
		// 拼接查询语句
		var sqlQuery = "SELECT " +
			sqlColConnect(col) +
			" FROM `" + prefix + table + "` " + extraString
		// 使用参数绑定，防止SQL注入
		rows, err := db.Query(sqlQuery)
		if err != nil {
			return nil, nil, err
		}
		// 获取查询结果字段集
		colArr, err := rows.Columns()
		if err != nil {
			return nil, nil, err
		}
		return rows, colArr, err
	}

}
