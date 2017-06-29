package handler

import (
	"github.com/ngonghi/VyLog/common"
)

type FileHandler struct {
	AbstractHandler
}

func (handler *FileHandler) Handle(message *common.Message) {
	if handler.IsHandling(message) {
		handler.Write(message)
	}
}

func (handler * FileHandler) Write(msg *common.Message) {
	buf := []byte{}
	buf = append(buf, handler.GetFormatter().Format(msg)...)
}