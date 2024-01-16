package sqli

import (
	dbConfig "Spark/server/config"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

func Connect() (*sql.DB, error) {

	// 检测数据库连接数据是否合法
	for k, v := range dbConfig.DBConfig {
		if v == "" {
			log.Printf("params error -> %s", k)
			return nil, errors.New(fmt.Sprintf("params error -> %s", k))
		}
	}
	// 拼凑数据库连接数据
	dbInfo := strings.Join([]string{dbConfig.DBConfig["user"], ":", dbConfig.DBConfig["pass"], "@tcp(",
		dbConfig.DBConfig["host"], ":", dbConfig.DBConfig["port"], ")/", dbConfig.DBConfig["database"], "?charset=",
		dbConfig.DBConfig["charset"]}, "")
	// 使用拼凑好的数据进行连接数据库
	dbIns, err := sql.Open(dbConfig.DBConfig["dbType"], dbInfo)

	if err != nil {
		return nil, err
	}
	if err = dbIns.Ping(); err != nil {
		return nil, err
	}
	// 设置每次连接声明周期以及最大连接闲置时间
	dbIns.SetConnMaxLifetime(100)
	dbIns.SetMaxIdleConns(10)
	return dbIns, nil

}
