package formatter

import "github.com/ngonghi/VyLog"

type Formatter interface {
	Format(message *VyLog.Message) string
}

type AbstractFormatter struct {
	logFormat string
	dateFormat string
}

func (f *AbstractFormatter) GetDefaultDateFormat() string {
	return "2016-08-31 02:17:00"
}

func (f *AbstractFormatter) GetDateFormat() string {
	if f.dateFormat == "" {
		return f.GetDefaultDateFormat()
	}

	return f.dateFormat
}