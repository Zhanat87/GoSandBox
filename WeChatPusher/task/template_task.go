package task

import (
	"github.com/hundredlee/wechat_pusher/enum"
	"github.com/hundredlee/wechat_pusher/models"
)

type TemplateTask struct {
	taskValue models.Message
}

func (self *TemplateTask) GetTaskType() string {
	return enum.TASK_TYPE_TEMPLATE
}

func (self *TemplateTask) SetTask(taskValue interface{}) {

	v, ok := taskValue.(models.Message)
	if ok {
		self.taskValue = v
	}
}

func (self *TemplateTask) GetTask() interface{} {
	return self.taskValue
}
