package service

import (
	"github.com/Deansquirrel/goMonitor/global"
	"github.com/Deansquirrel/goMonitor/notify"
	"os"
	"os/signal"
	"syscall"
)
import log "github.com/Deansquirrel/goToolLog"

const (
	//测试WebHookKey
	TestWebHookKey = "7a84d09b83f9633ad37866505d2c0c26e39f4fa916b3af8f6a702180d3b9906b"
)

type cronService struct {
}

func NewCronService() *cronService {
	return &cronService{}
}

func (cs *cronService) Start() {
	cs.start()
	select {
	case <-global.Ctx.Done():
		return
	}
}

func (cs *cronService) start() {
	log.Debug("CronService starting")
	defer log.Debug("CronService start complete")
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch,
			os.Interrupt,
			syscall.SIGINT,
			os.Kill,
			syscall.SIGKILL,
			syscall.SIGTERM,
		)
		select {
		case <-ch:
			defer global.Cancel()
		case <-global.Ctx.Done():
		}
	}()
}

//测试消息发送
func (cs *cronService) test() {
	dt := notify.NewDingTalkRobot(global.SysConfig.DingTalkConfig.Address)
	var err error
	err = dt.SendTextMsg(TestWebHookKey, "normal msg")
	if err != nil {
		log.Debug("ERR:" + err.Error())
	}
	atList := make([]string, 0)
	atList = append(atList, "15298386821")
	err = dt.SendTextMsgWithAtList(TestWebHookKey, "at msg", atList)
	if err != nil {
		log.Debug(err.Error())
	}
	err = dt.SendTextMsgWithAtAll(TestWebHookKey, "at all msg")
	if err != nil {
		log.Debug(err.Error())
	}
}
