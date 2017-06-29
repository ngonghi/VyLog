package common

import "time"

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
