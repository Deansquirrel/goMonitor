package taskConfigRepository

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitor/global"
	"github.com/Deansquirrel/goToolMSSql"
)

//获取配置库连接配置
func getConfigDBConfig() *goToolMSSql.MSSqlConfig {
	return &goToolMSSql.MSSqlConfig{
		Server: global.SysConfig.ConfigDBConfig.Server,
		Port:   global.SysConfig.ConfigDBConfig.Port,
		DbName: global.SysConfig.ConfigDBConfig.DbName,
		User:   global.SysConfig.ConfigDBConfig.User,
		Pwd:    global.SysConfig.ConfigDBConfig.Pwd,
	}
}

func getRowsBySQL(sql string, args ...interface{}) (*sql.Rows, error) {
	conn, err := goToolMSSql.GetConn(getConfigDBConfig())
	if err != nil {
		return nil, err
	}
	if args == nil {
		rows, err := conn.Query(sql)
		if err != nil {
			return nil, err
		}
		return rows, nil
	} else {
		rows, err := conn.Query(sql, args)
		if err != nil {
			return nil, err
		}
		return rows, nil
	}
}
