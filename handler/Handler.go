package handler

import (
	vFormatter "github.com/ngonghi/VyLog/formatter"
	"github.com/ngonghi/VyLog/common"
)

type Handler interface {
	IsHandling(message *common.Message) bool

	Handle(message *common.Message) error

	SetFormatter(formatter vFormatter.Formatter)

	GetFormatter() vFormatter.Formatter
}

type AbstractHandler struct {
	Level     int
	Formatter vFormatter.Formatter
}

func (handler *AbstractHandler) GetLevel() int {
	return handler.Level
}

func (handler *AbstractHandler) SetLevel(level int) {
	handler.Level = level
}

func (handler *AbstractHandler) IsHandling(msg *common.Message) bool {
	return msg.GetLevel() >= handler.GetLevel()
}

func (handler *AbstractHandler) SetFormatter(formatter vFormatter.Formatter) {
	handler.Formatter = formatter
}

func (handler *AbstractHandler) GetFormatter() vFormatter.Formatter {

	if handler.Formatter == nil {
		return handler.GetDefaultFormatter()
	}

	return handler.Formatter
}

func (handler *AbstractHandler) GetDefaultFormatter() vFormatter.LineFormatter {
	return vFormatter.LineFormatter{}
}


