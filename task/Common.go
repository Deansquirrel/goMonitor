package task

import (
	"github.com/Deansquirrel/goMonitor/taskConfigRepository"
	"github.com/robfig/cron"
)

var intTaskConfigRepository taskConfigRepository.IntTaskConfig
var intConfigList map[string]*taskConfigRepository.IntTaskConfigData
var intTaskList map[string]*cron.Cron

func init() {
	intTaskConfigRepository = taskConfigRepository.IntTaskConfig{}
	intConfigList = make(map[string]*taskConfigRepository.IntTaskConfigData)
	intTaskList = make(map[string]*cron.Cron)
}
