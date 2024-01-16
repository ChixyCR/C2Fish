package sqli

import (
	dbConfig "Spark/server/config"
	"database/sql"
	"errors"
	"fmt"
)

type Sql struct {
	This interface{}
}

// verify
func dataVerify(i interface{}) string {
	switch i.(type) {
	case []uint8:
		return "string"
	case int64:
		return "int64"
	default:
		return "error"
	}
}

// connect
func (s *Sql) Connect() (*sql.DB, error) {
	sqlIns, err := Connect()
	if err != nil {
		return nil, err
	}
	s.This = sqlIns
	return (s.This).(*sql.DB), nil
}

// get
func (s *Sql) Self() (*sql.DB, error) {
	sqlIns, err := (s.This).(*sql.DB)
	if !err {
		return nil, errors.New(fmt.Sprint("type error -> undefined"))
	}
	return sqlIns, nil
}

// query
func (s *Sql) Query(tableTag string, col []string, condCol []string, condVal []string, pad string, extraString string) ([]map[string]interface{}, error) {

	sqlIns, err := s.Self()
	if err != nil {
		return nil, err
	}

	tablePrefix := dbConfig.DBConfig["tablePrefix"]
	tableName := dbConfig.TableConfig[tableTag]["tableName"]

	rows, colArr, err := QuerySafe(sqlIns, tableName, tablePrefix, col, condCol, condVal, pad, extraString)
	if err != nil {
		return nil, err
	}

	tableData := make([]map[string]interface{}, 0)
	for rows.Next() {

		dataInterface := make([]interface{}, len(colArr))
		dataMap := make(map[string]interface{}, len(colArr))

		for index, _ := range dataInterface {
			dataInterface[index] = new(interface{})
		}

		err = rows.Scan(dataInterface...)
		if err != nil {
			return nil, err
		}

		for index, colName := range colArr {
			switch dataVerify(*dataInterface[index].(*interface{})) {
			case "string":
				dataMap[colName] = string((*dataInterface[index].(*interface{})).([]uint8))
			case "int64":
				dataMap[colName] = (*dataInterface[index].(*interface{})).(int64)
			case "error":
				return nil, errors.New(fmt.Sprint("type error -> unknown data type"))
			}
		}

		tableData = append(tableData, dataMap)
	}

	return tableData, nil
}

// update
func (s *Sql) Update(tableTag string, upCol []string, upVal []string, condCol []string, condVal []string, pad string, extraString string) (int64, error) {

	sqlIns, err := s.Self()
	if err != nil {
		return 0, err
	}

	tablePrefix := dbConfig.DBConfig["tablePrefix"]
	tableName := dbConfig.TableConfig[tableTag]["tableName"]

	return UpdateQuery(sqlIns, tableName, tablePrefix, upCol, upVal, condCol, condVal, pad, extraString)
}

// insert
func (s *Sql) Insert(tableTag string, insCol []string, insVal []string, extraString string) (int64, error) {
	sqlIns, err := s.Self()
	if err != nil {
		return 0, err
	}

	tablePrefix := dbConfig.DBConfig["tablePrefix"]
	tableName := dbConfig.TableConfig[tableTag]["tableName"]

	return InsertSafe(sqlIns, tableName, tablePrefix, insCol, insVal, extraString)

}

// delete
func (s *Sql) Delete(tableTag string, delCol []string, delVal []string, pad string, extraString string) (int64, error) {

	sqlIns, err := s.Self()
	if err != nil {
		return 0, err
	}

	tablePrefix := dbConfig.DBConfig["tablePrefix"]
	tableName := dbConfig.TableConfig[tableTag]["tableName"]

	return DeleteSafe(sqlIns, tableName, tablePrefix, delCol, delVal, pad, extraString)

}

// close
func (s *Sql) Close() error {
	sqlIns, err := s.Self()
	if err != nil {
		return err
	}
	return sqlIns.Close()
}
