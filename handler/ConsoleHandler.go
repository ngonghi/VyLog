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

func (handler *ConsoleHandler) Handle(message *common.Message) error {
	if handler.IsHandling(message) {
		err := handler.Write(message)
		return err
	}

	return nil
}

func (handler * ConsoleHandler) Write(msg *common.Message) error {
	buf := []byte{}
	buf = append(buf, handler.GetFormatter().Format(msg)...)
	_, err := handler.GetOutput().Write(buf)

	return err
}

func (handler * ConsoleHandler) GetOutput() io.WriteCloser {

	if handler.out != nil {
		return handler.out
	}

	return os.Stdout
}