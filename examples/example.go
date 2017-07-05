package main

import (
	"github.com/ngonghi/VyLog/handler"
	"github.com/ngonghi/VyLog"
)

func main()  {
	consoleHandler := handler.GetConsoleHandler()
	fileHandler,err := handler.GetFileHandler("test", 0)
	//slackHandler := handler.GetSlackHandler("#test", "https://hook", "VyLog", "TEST")

	v := &VyLog.Vylog{}
	v.AddHandler(consoleHandler)

	if err == nil {
		v.AddHandler(fileHandler)
	}

	//v.AddHandler(slackHandler)

	v.Trace("TEST")
	v.Debug("TEST")
	v.Info("TEST")
	v.Warn("TEST")
	v.Error("TEST")
	v.Fatal("TEST")
}