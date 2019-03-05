package task

import (
	"github.com/Deansquirrel/goMonitor/global"
	"github.com/Deansquirrel/goMonitor/object"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goToolMSSql"
)

const SqlGetTaskMConfig = "" +
	"SELECT [FID],[FTitle],[FRemark] FROM [MConfig]"

func GetTaskMConfig() ([]object.TaskMConfig, error) {
	conn, err := goToolMSSql.GetConn(getConfigDBConfig())
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(SqlGetTaskMConfig)
	if err != nil {
		return nil, err
	}
	defer func() {
		errLs := rows.Close()
		if errLs != nil {
			log.Error(errLs.Error())
		}
	}()
	//TODO 获取rows中的数据
	return nil, nil
}

func getConfigDBConfig() *goToolMSSql.MSSqlConfig {
	return &goToolMSSql.MSSqlConfig{
		Server: global.SysConfig.ConfigDBConfig.Server,
		Port:   global.SysConfig.ConfigDBConfig.Port,
		DbName: global.SysConfig.ConfigDBConfig.DbName,
		User:   global.SysConfig.ConfigDBConfig.User,
		Pwd:    global.SysConfig.ConfigDBConfig.Pwd,
	}
}
