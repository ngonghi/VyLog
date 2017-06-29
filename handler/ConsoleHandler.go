package handler

import (
	"github.com/ngonghi/VyLog"
	"io"
	"os"
)

type ConsoleHandler struct {
	AbstractHandler

	out        io.WriteCloser
}

func (handler *ConsoleHandler) Handle(message *VyLog.Message) {
	if handler.IsHandling(message) {
		handler.Write(message)
	}
}

func (handler * ConsoleHandler) Write(msg * VyLog.Message) {
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