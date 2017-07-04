package handler

import (
	"github.com/ngonghi/VyLog/common"
	"github.com/ngonghi/VyLog"
	"encoding/json"
)

const (
	SLACK_COLOR_DANGER = "danger"
	SLACK_COLOR_WARNING = "warning"
	SLACK_COLOR_GOOD = "good"
	SLACK_COLOR_DEFAULT = "#e3e4e6"
)

type SlackHandler struct {
	AbstractHandler

	channel string
	hookUrl string
	username string
	preText string
}

func (handler *SlackHandler) getColor(level int) string {
	switch level {
	case VyLog.INFO:
		return SLACK_COLOR_GOOD
	case VyLog.WARN:
		return SLACK_COLOR_WARNING
	case VyLog.ERROR, VyLog.FATAL:
		return SLACK_COLOR_DANGER
	default:
		return SLACK_COLOR_DEFAULT
	}
}

func (handler *SlackHandler) Handle(message *common.Message) {
	if handler.IsHandling(message) {
		handler.Post(message)
	}
}

func (handler * SlackHandler) Post(msg *common.Message) {
	buf := []byte{}
	buf = append(buf, handler.GetFormatter().Format(msg)...)
}

//func (handler * SlackHandler) createSlackMsgContent(msg *common.Message) string {
//
//
//
//
//}