package handler

import (
	"github.com/ngonghi/VyLog/common"
)

type SlackHandler struct {
	AbstractHandler
}

func (handler *SlackHandler) Handle(message *common.Message) {
	if handler.IsHandling(message) {
		handler.Write(message)
	}
}

func (handler * SlackHandler) Write(msg *common.Message) {
	buf := []byte{}
	buf = append(buf, handler.GetFormatter().Format(msg)...)
}