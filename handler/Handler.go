package handler

import (
	vFormatter "github.com/ngonghi/VyLog/formatter"
	"github.com/ngonghi/VyLog/common"
)

type Handler interface {
	IsHandling(message *common.Message) bool

	Handle(message *common.Message)

	SetFormatter(formatter vFormatter.Formatter)

	GetFormatter() vFormatter.Formatter
}

type AbstractHandler struct {
	level     int
	formatter vFormatter.Formatter
}

func (handler *AbstractHandler) GetLevel() int {
	return handler.level
}

func (handler *AbstractHandler) IsHandling(msg *common.Message) bool {
	return msg.GetLevel() >= handler.GetLevel()
}

func (handler *AbstractHandler) SetFormatter(formatter vFormatter.Formatter) {
	handler.formatter = formatter
}

func (handler *AbstractHandler) GetFormatter() vFormatter.Formatter {

	if handler.formatter == nil {
		return handler.GetDefaultFormatter()
	}

	return handler.formatter
}

func (handler *AbstractHandler) GetDefaultFormatter() vFormatter.LineFormatter {
	return vFormatter.LineFormatter{}
}


