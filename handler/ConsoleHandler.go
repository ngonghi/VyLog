package handler

import (
	"io"
	"os"
	"github.com/ngonghi/VyLog/common"
)

type ConsoleHandler struct {
	AbstractHandler

	out        io.WriteCloser
}

func GetConsoleHandler() *ConsoleHandler {
	return &ConsoleHandler{}
}

func (handler *ConsoleHandler) Handle(message *common.Message) {
	if handler.IsHandling(message) {
		handler.Write(message)
	}
}

func (handler * ConsoleHandler) Write(msg *common.Message) {
	buf := []byte{}
	buf = append(buf, handler.GetFormatter().Format(msg)...)
	handler.GetOutput().Write(buf)
}

func (handler * ConsoleHandler) GetOutput() io.WriteCloser {

	if handler.out != nil {
		return handler.out
	}

	return os.Stdout
}