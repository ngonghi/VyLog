package VyLog

import (
	"time"
	"github.com/ngonghi/VyLog/handler"
	"github.com/ngonghi/VyLog/common"
	"runtime"
	"fmt"
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
	case common.TRACE:
		return "TRACE"
	case common.DEBUG:
		return "DEBUG"
	case common.INFO:
		return "INFO"
	case common.WARN:
		return "WARN"
	case common.ERROR:
		return "ERROR"
	case common.FATAL:
		return "FATAL"
	}

	return "UNKNOWN"
}

func GetLevelByName(levelName string) int {

	switch levelName {
	case "TRACE":
		return common.TRACE
	case "DEBUG":
		return common.DEBUG
	case "INFO":
		return common.INFO
	case "WARN":
		return common.WARN
	case "ERROR":
		return common.ERROR
	case "FATAL":
		return common.FATAL
	}
	
	return -1
}

func (log *Vylog) Trace(msg string) {

	pc, fn, line, _ := runtime.Caller(1)
	traceMsg := fmt.Sprintf("%s[%s:%d]%s",runtime.FuncForPC(pc).Name(), fn, line, msg)

	log.log(common.TRACE, traceMsg)
}

func (log *Vylog) Debug(msg string) {

	pc, fn, line, _ := runtime.Caller(1)
	debugMsg := fmt.Sprintf("%s[%s:%d]%s",runtime.FuncForPC(pc).Name(), fn, line, msg)

	log.log(common.DEBUG, debugMsg)
}

func (log *Vylog) Info(msg string) {
	log.log(common.INFO, msg)
}

func (log *Vylog) Warn(msg string) {
	log.log(common.WARN, msg)
}

func (log *Vylog) Error(msg string) {

	pc, fn, line, _ := runtime.Caller(1)
	errorMsg := fmt.Sprintf("%s[%s:%d]%s",runtime.FuncForPC(pc).Name(), fn, line, msg)

	log.log(common.ERROR, errorMsg)
}

func (log *Vylog) Fatal(msg string) {

	pc, fn, line, _ := runtime.Caller(1)
	fatalMsg := fmt.Sprintf("%s[%s:%d]%s",runtime.FuncForPC(pc).Name(), fn, line, msg)

	log.log(common.FATAL, fatalMsg)
}