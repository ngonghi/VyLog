package formatter

import (
	"github.com/ngonghi/VyLog"

	"fmt"
)

type LineFormatter struct {
	AbstractFormatter
}

func (lf LineFormatter) Format(message *VyLog.Message) string {
	return fmt.Sprintf(lf.getLogFormat(), message.GetTime().Format(lf.GetDateFormat()), message.GetLevelName(), message.GetMessage())
}

func (lf *LineFormatter) setLogFormat(logFormat string) {
	lf.logFormat = logFormat
}

func (lf *LineFormatter) getLogFormat() string {
	if lf.logFormat == "" {
		return lf.getDefaultLogFormat()
	}

	return lf.logFormat
}

func (lf *LineFormatter) getDefaultLogFormat() string {
	return "[%s] %s: %s\n"
}