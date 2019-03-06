package taskConfigRepository

import (
	"github.com/Deansquirrel/goToolCommon"
	"strings"
)

type IntTaskConfig struct {
	FId         string
	FServer     string
	FPort       int
	FDbName     string
	FDbUser     string
	FDbPwd      string
	FSearch     string
	FCron       string
	FCheckMax   int
	FCheckMin   int
	FMsgTitle   string
	FMsgContent string
}

func NewIntTaskConfig(server string, port int, dbName string, dbUser string, dbPwd string,
	search string, cron string, checkMax int, checkMin int, msgTitle string, msgContent string) *IntTaskConfig {
	return &IntTaskConfig{
		FId:         strings.ToUpper(goToolCommon.Guid()),
		FServer:     server,
		FPort:       port,
		FDbName:     dbName,
		FDbUser:     dbUser,
		FDbPwd:      dbPwd,
		FSearch:     search,
		FCron:       cron,
		FCheckMax:   checkMax,
		FCheckMin:   checkMin,
		FMsgTitle:   msgTitle,
		FMsgContent: msgContent,
	}
}
