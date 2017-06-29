package main

import (
	"github.com/ngonghi/VyLog/handler"
	"github.com/ngonghi/VyLog"
)

func main()  {
	l := &handler.ConsoleHandler{}
	v := &VyLog.Vylog{}
	v.AddHandler(l)
	
	v.Trace("TEST")
	v.Debug("TEST")
	v.Info("TEST")
	v.Warn("TEST")
	v.Error("TEST")
	v.Fatal("TEST")
}