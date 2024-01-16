package sqli

import (
	lib "Spark/server/lib/func"
	"database/sql"
)

func DeleteSafe(db *sql.DB, table string, prefix string, delCol []string, delVal []string, pad string, extraString string) (int64, error) {
	var sqlQuery = "DELETE FROM `" + prefix + table + "` WHERE " + sqlDeleteConnect(delCol, pad) + extraString
	stmt, err := db.Prepare(sqlQuery)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(lib.MergeSliceStringValue(delVal)...)

	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
