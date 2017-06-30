package formatter

import (
	"github.com/ngonghi/VyLog/common"
)

type Formatter interface {
	Format(message *common.Message) string
}

type AbstractFormatter struct {
	logFormat string
	dateFormat string
}

func (f *AbstractFormatter) GetDefaultDateFormat() string {
	return "2006-01-02 15:04:05"
}

func (f *AbstractFormatter) GetDateFormat() string {
	if f.dateFormat == "" {
		return f.GetDefaultDateFormat()
	}

	return f.dateFormat
}