package service

import "github.com/Deansquirrel/goMonitor/task"

var intTask task.IntTask

func init() {
	intTask = task.IntTask{}
}
