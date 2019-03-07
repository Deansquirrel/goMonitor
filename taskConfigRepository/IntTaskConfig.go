package taskConfigRepository

import (
	"database/sql"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"strings"
)

const SqlGetIntTaskConfig = "" +
	"SELECT [FId],[FServer],[FPort],[FDbName],[FDbUser]," +
	"[FDbPwd],[FSearch],[FCron],[FCheckMax],[FCheckMin]," +
	"[FMsgTitle],[FMsgContent] " +
	"FROM [IntTaskConfig]"

const SqlGetIntTaskConfigById = "" +
	"SELECT [FId],[FServer],[FPort],[FDbName],[FDbUser]," +
	"[FDbPwd],[FSearch],[FCron],[FCheckMax],[FCheckMin]," +
	"[FMsgTitle],[FMsgContent] " +
	"FROM [IntTaskConfig] " +
	"WHERE [FId]=?"

type IntTaskConfig struct {
}

type intTaskConfigData struct {
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
	search string, cron string, checkMax int, checkMin int, msgTitle string, msgContent string) *intTaskConfigData {
	return &intTaskConfigData{
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

func (itc *IntTaskConfig) GetIntTaskConfigList() ([]intTaskConfigData, error) {
	rows, err := comm.getRowsBySQL(SqlGetIntTaskConfig)
	if err != nil {
		return nil, err
	}
	return itc.getIntTaskConfigListByRows(rows)
}

func (itc *IntTaskConfig) GetIntTaskConfig(id string) ([]intTaskConfigData, error) {
	rows, err := comm.getRowsBySQL(SqlGetIntTaskConfigById, id)
	if err != nil {
		return nil, err
	}
	return itc.getIntTaskConfigListByRows(rows)
}

func (itc *IntTaskConfig) getIntTaskConfigListByRows(rows *sql.Rows) ([]intTaskConfigData, error) {
	defer func() {
		errLs := rows.Close()
		if errLs != nil {
			log.Error(errLs.Error())
		}
	}()
	var fId, fServer, fDbName, fDbUser, fDbPwd, fSearch, fCron, fMsgTitle, fMsgContent string
	var fPort, fCheckMax, fCheckMin int
	resultList := make([]intTaskConfigData, 0)
	for rows.Next() {
		err := rows.Scan(
			&fId, &fServer, &fPort, &fDbName, &fDbUser,
			&fDbPwd, &fSearch, &fCron, &fCheckMax, &fCheckMin,
			&fMsgTitle, &fMsgContent)
		if err != nil {
			return nil, err
		}
		config := intTaskConfigData{
			FId:         fId,
			FServer:     fServer,
			FPort:       fPort,
			FDbName:     fDbName,
			FDbUser:     fDbUser,
			FDbPwd:      fDbPwd,
			FSearch:     fSearch,
			FCron:       fCron,
			FCheckMax:   fCheckMax,
			FCheckMin:   fCheckMin,
			FMsgTitle:   fMsgTitle,
			FMsgContent: fMsgContent,
		}
		resultList = append(resultList, config)
	}
	return resultList, nil
}
