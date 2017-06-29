package handler

import (
	"github.com/ngonghi/VyLog/common"
	"os"
	"fmt"
)

const (
	APPEND = iota
	ROTATE
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

	maxRotateFile int
	mustRotate bool
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

func (handler *FileHandler) GetFileLogger(filePath string, mode int, maxRotateFile int) (*FileHandler, error) {
	var file *os.File
	var err error

	switch mode {
	case APPEND:
		file, err = os.OpenFile(f, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	case ROTATE:
		return nil, fmt.Errorf("Invalid Log File Mode: %d", mode)
	default:
		return nil, fmt.Errorf("Can Open Log In This Mode: %d", mode)
	}
	
	if err != nil {
		return nil, fmt.Errorf("Open log file error: %s", err)
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
}
