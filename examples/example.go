package main

import (
	"github.com/ngonghi/VyLog"
	"time"

	"github.com/ngonghi/VyLog/handler"
)

func main()  {
	m := &VyLog.Message {Level : 1, LevelName : "tEST", Message : "test", Time : time.Now()}
	l := &handler.ConsoleHandler{}
	l.Handle(m)
}