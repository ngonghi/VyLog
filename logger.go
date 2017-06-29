package VyLog

import "time"

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

type LoggerInterface interface {

	Trace(string, ...interface{})

	Debug(string, ...interface{})

	Info(string, ...interface{})

	Warn(string, ...interface{})

	Error(string, ...interface{})

	Fatal(string, ...interface{})

	Log(int, string, ...interface{})
}

type Message struct {
	Level int
	LevelName string
	Message string
	Time time.Time
}

func (msg *Message) GetLevel() int {
	return msg.Level
}

func (msg *Message) GetLevelName() string {
	return msg.LevelName
}

func (msg * Message) GetMessage() string  {
	return msg.Message
}

func (msg * Message) GetTime() time.Time {
	return msg.Time
}