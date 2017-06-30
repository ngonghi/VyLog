package handler

import (
	"github.com/ngonghi/VyLog/common"
	"os"
	"fmt"
)

const (
	APPEND = iota
)

const (
	BUFSIZE = 100
)

type FileHandler struct {
	AbstractHandler

	msgQueue chan *common.Message

	mode     int
	file     *os.File
	filePath string
}

func GetFileHandler(filePath string, mode int) (*FileHandler, error) {

	var file *os.File
	var err error

	switch mode {
	case APPEND:
		file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	default:
		return nil, fmt.Errorf("Can Open Log In This Mode: %d", mode)
	}

	if err != nil {
		return nil, fmt.Errorf("Open log file error: %s", err)
	}

	return newHandler(file), nil
}

func (handler *FileHandler) SetMode(mode int) {
	handler.mode = mode
}

func (handler *FileHandler) GetMode() int {
	return handler.mode
}

func (handler *FileHandler) GetFilePath() string {
	return handler.filePath
}

func (handler *FileHandler) SetFilePath(filePath string) {
	handler.filePath = filePath
}

func newHandler(file *os.File) (h *FileHandler) {
	h = &FileHandler{
		msgQueue: make(chan *common.Message, BUFSIZE),
		file: file,
	}
	return
}

func (handler *FileHandler) startOutput() {
	for {
		m, ok := <-handler.msgQueue
		if !ok {
			handler.file.Sync()
			return
		}
		handler.Write(m)
	}
}

func (handler *FileHandler) Handle(message *common.Message) {
	if handler.IsHandling(message) {
		handler.Write(message)
	}
}

func (handler *FileHandler) Write(msg *common.Message) {
	buf := []byte{}
	buf = append(buf, handler.GetFormatter().Format(msg)...)
	handler.file.Write(buf)
}