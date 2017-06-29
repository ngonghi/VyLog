package VyLog

import (
	"time"
	"github.com/ngonghi/VyLog/handler"
	"github.com/ngonghi/VyLog/common"
)

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

type LoggerInterface interface {

	Trace(string)

	Debug(string)

	Info(string)

	Warn(string)

	Error(string)

	Fatal(string)

	log(int, string)
}

type Vylog struct {
	handlers []handler.Handler
}

func (log *Vylog) AddHandler(handlers ...handler.Handler) {
	for _, handler := range handlers {
		log.handlers = append(log.handlers, handler)
	}
}

func (log *Vylog) log(level int, msg string) {
	for _, handler := range log.handlers {
		handler.Handle(log.createMessage(level, msg))
	}
}

func (log *Vylog) createMessage(level int, msg string) *common.Message {
	return &common.Message{
		Level: level,
		LevelName: log.getLevelName(level),
		Message: msg,
		Time: time.Now(),
	}
}

func (log *Vylog) getLevelName(level int) string {
	
	switch level {
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}

	return "UNKNOWN"
}

func (log *Vylog) Trace(msg string) {
	log.log(TRACE, msg)
}

func (log *Vylog) Debug(msg string) {
	log.log(DEBUG, msg)
}

func (log *Vylog) Info(msg string) {
	log.log(INFO, msg)
}

func (log *Vylog) Warn(msg string) {
	log.log(WARN, msg)
}

func (log *Vylog) Error(msg string) {
	log.log(ERROR, msg)
}

func (log *Vylog) Fatal(msg string) {
	log.log(FATAL, msg)
}