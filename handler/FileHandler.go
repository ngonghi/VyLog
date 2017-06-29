package handler

import "github.com/ngonghi/VyLog"

type FileHandler struct {
	AbstractHandler
}

func (handler *FileHandler) Handle(message *VyLog.Message) {
	if handler.IsHandling(message) {
		handler.Write(message)
	}
}

func (handler * FileHandler) Write(msg * VyLog.Message) {
	buf := []byte{}
	buf = append(buf, handler.GetFormatter().Format(msg)...)
}