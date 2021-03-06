package task

import (
	"fmt"
	"github.com/Deansquirrel/goMonitor/taskConfigRepository"
	"github.com/Deansquirrel/goMonitor/worker"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/kataras/iris/core/errors"
	"github.com/robfig/cron"
)

type IntTask struct {
}

func (it *IntTask) StartTask() error {
	//获取配置列表
	list, err := intTaskConfigRepository.GetIntTaskConfigList()
	if err != nil {
		return err
	}
	//清空Config列表
	intConfigList = make(map[string]*taskConfigRepository.IntTaskConfigData)
	//清空Task列表
	intTaskList = make(map[string]*cron.Cron)
	//缓存配置列表、任务列表
	errMsg := ""
	errMsgFormat := "添加任务[%s]报错：%s；"
	for _, config := range list {
		err = it.addTask(config)
		if err != nil {
			errMsg = errMsg + fmt.Sprintf(errMsgFormat, config.FId, err.Error())
		}
	}
	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}

func (it *IntTask) StartJob(id string) error {
	c := intTaskList[id]
	if c == nil {
		return errors.New("Task is not exist")
	}
	c.Start()
	return nil
}

func (it *IntTask) StopJob(id string) error {
	c := intTaskList[id]
	if c == nil {
		return errors.New("Task is not exist")
	}
	c.Stop()
	return nil
}

func (it *IntTask) RefreshConfig() error {
	//获取配置列表
	list, err := intTaskConfigRepository.GetIntTaskConfigList()
	if err != nil {
		return err
	}
	listId := make([]string, 0)
	mapId := make(map[string]*taskConfigRepository.IntTaskConfigData, 0)
	for _, config := range list {
		listId = append(listId, config.FId)
		mapId[config.FId] = config
	}

	configId := make([]string, 0)
	for key := range intConfigList {
		configId = append(configId, key)
	}

	addList, delList, checkList := goToolCommon.CheckDiff(listId, configId)

	errMsg := ""
	errMsgFormat := "刷新任务[%s]报错：%s；"

	for _, id := range addList {
		err = it.addTask(mapId[id])
		if err != nil {
			errMsg = errMsg + fmt.Sprintf(errMsgFormat, id, err.Error())
		}
	}
	for _, id := range delList {
		it.removeTask(id)
	}
	for _, id := range checkList {
		err = it.checkTask(mapId[id])
		if err != nil {
			errMsg = errMsg + fmt.Sprintf(errMsgFormat, id, err.Error())
		}
	}
	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}

func (it *IntTask) addTask(config *taskConfigRepository.IntTaskConfigData) error {
	//------------------------------------------------------------------------------------------------------------------
	configStr, err := goToolCommon.GetJsonStr(config)
	if err != nil {
		log.Warn(fmt.Sprintf("Add Int Task，转换配置内容时遇到错误:%s，configID：%s", configStr, config.FId))
	} else {
		log.Warn(fmt.Sprintf("Add Int Task:%s", configStr))
	}
	//------------------------------------------------------------------------------------------------------------------
	intConfigList[config.FId] = config
	w := worker.NewIntWorker(config)
	c := cron.New()
	err = c.AddJob(config.FCron, w)
	if err != nil {
		log.Error(err.Error())
	} else {
		c.Start()
	}
	intTaskList[config.FId] = c
	return err
}

func (it *IntTask) checkTask(config *taskConfigRepository.IntTaskConfigData) error {
	exConfig := intConfigList[config.FId]
	if exConfig == nil {
		return it.addTask(config)
	}
	if exConfig.IsEqual(config) {
		return nil
	}
	it.removeTask(config.FId)
	return it.addTask(config)
}

func (it *IntTask) removeTask(id string) {
	config := intConfigList[id]
	if config == nil {
		log.Warn(fmt.Sprintf("remove task :config is not exist,taskId[%s]", id))
	} else {
		configStr, err := goToolCommon.GetJsonStr(config)
		if err != nil {
			log.Warn(fmt.Sprintf("Del Int Task，转换配置内容时遇到错误:%s，configID：%s", configStr, config.FId))
		} else {
			log.Warn(fmt.Sprintf("Del Int Task:%s", configStr))
		}
		delete(intConfigList, id)
	}
	c := intTaskList[id]
	if c == nil {
		log.Warn(fmt.Sprintf("remove task :task is not exist,taskId[%s]", id))
	} else {
		c.Stop()
		delete(intTaskList, id)
	}
}
