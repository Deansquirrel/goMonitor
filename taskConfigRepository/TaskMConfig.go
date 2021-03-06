package taskConfigRepository

import (
	"database/sql"
	log "github.com/Deansquirrel/goToolLog"
)

const SqlGetTaskMConfig = "" +
	"SELECT [FID],[FTitle],[FRemark] " +
	"FROM [MConfig]"
const SqlGetTaskMConfigById = "" +
	"SELECT [FID],[FTitle],[FRemark] " +
	"FROM [MConfig] WHERE [FID] = ?"

type TaskMConfig struct {
}

type taskMConfigData struct {
	FId     string
	FTitle  string
	FRemark string
}

//func NewTaskMConfig(title string, remark string) *taskMConfigData {
//	return &taskMConfigData{
//		FId:     strings.ToUpper(goToolCommon.Guid()),
//		FTitle:  title,
//		FRemark: remark,
//	}
//}

func (tmc *TaskMConfig) GetMConfigList() ([]*taskMConfigData, error) {
	rows, err := comm.getRowsBySQL(SqlGetTaskMConfig)
	if err != nil {
		return nil, err
	}
	return tmc.getMConfigListByRows(rows)
}

func (tmc *TaskMConfig) GetMConfig(id string) ([]*taskMConfigData, error) {
	rows, err := comm.getRowsBySQL(SqlGetTaskMConfigById, id)
	if err != nil {
		return nil, err
	}
	return tmc.getMConfigListByRows(rows)
}

func (tmc *TaskMConfig) getMConfigListByRows(rows *sql.Rows) ([]*taskMConfigData, error) {
	defer func() {
		errLs := rows.Close()
		if errLs != nil {
			log.Error(errLs.Error())
		}
	}()
	var fId, fTitle, fRemark string
	resultList := make([]*taskMConfigData, 0)
	for rows.Next() {
		err := rows.Scan(&fId, &fTitle, &fRemark)
		if err != nil {
			return nil, err
		}
		config := taskMConfigData{
			FId:     fId,
			FTitle:  fTitle,
			FRemark: fRemark,
		}
		resultList = append(resultList, &config)
	}
	return resultList, nil
}
