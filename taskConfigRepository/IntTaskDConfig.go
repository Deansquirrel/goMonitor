package taskConfigRepository

import (
	"database/sql"
	log "github.com/Deansquirrel/goToolLog"
)

//const SqlGetIntTaskDConfig = "" +
//	"SELECT [FID],[FMsgSearch] " +
//	"FROM [IntTaskDConfig]"

const SqlGetIntTaskDConfigById = "" +
	"SELECT [FID],[FMsgSearch] " +
	"FROM [IntTaskDConfig] " +
	"WHERE [FId]=?"

type IntTaskDConfig struct {
}

type intTaskDConfigData struct {
	FId        string
	FMsgSearch string
}

//func (itc *IntTaskDConfig) GetIntTaskDConfigList() ([]intTaskDConfigData, error) {
//	rows, err := comm.getRowsBySQL(SqlGetIntTaskDConfig)
//	if err != nil {
//		return nil, err
//	}
//	return itc.getIntTaskDConfigListByRows(rows)
//}

func (itc *IntTaskDConfig) GetIntTaskDConfig(id string) ([]*intTaskDConfigData, error) {
	rows, err := comm.getRowsBySQL(SqlGetIntTaskDConfigById, id)
	if err != nil {
		return nil, err
	}
	return itc.getIntTaskDConfigListByRows(rows)
}

func (itc *IntTaskDConfig) getIntTaskDConfigListByRows(rows *sql.Rows) ([]*intTaskDConfigData, error) {
	defer func() {
		errLs := rows.Close()
		if errLs != nil {
			log.Error(errLs.Error())
		}
	}()
	var fId, fMsgSearch string
	resultList := make([]*intTaskDConfigData, 0)
	for rows.Next() {
		err := rows.Scan(&fId, &fMsgSearch)
		if err != nil {
			return nil, err
		}
		config := intTaskDConfigData{
			FId:        fId,
			FMsgSearch: fMsgSearch,
		}
		resultList = append(resultList, &config)
	}
	return resultList, nil
}
