package object

import (
	"github.com/Deansquirrel/goToolCommon"
	"strings"
)

type TaskMConfig struct {
	FId     string
	FTitle  string
	FRemark string
}

func NewTaskConfig(title string, remark string) *TaskMConfig {
	return &TaskMConfig{
		FId:     strings.ToUpper(goToolCommon.Guid()),
		FTitle:  title,
		FRemark: remark,
	}
}
