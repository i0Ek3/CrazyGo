package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// please ref https://beego.me/docs/mvc/controller/logs.md

func main() {
	l := logs.GetLogger()
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(true)
	logs.SetLogger(logs.AdapterConsole, `{"level":8, "color":true}`) // output to console with color support, default output to os.Stdout
	//logs.SetLogger(logs.AdapterFile, `{"filename":"test.log"}`)
	beego.SetLogger("file", `{"filename":"./test.log"}`)
	logs.EnableFuncCallDepth(true) // show you issue file
	//logs.Async(1e3) // async output logs
	l.Println("this is a msg of HTTP")
	logs.GetLogger("HTTP").Println("this is a msg of HTTP")
	logs.GetLogger("ORM").Println("this is a msg of ORM")

	/*
		logs.Debug("My first time to use Go is ", 2017)
		logs.Info("Go have %v years old made by %s", 10, "Google")
		logs.Warn("JSON is a type of K-V like", map[string]int{"key": 2018})
		logs.Error(1024, "is a very", "good game")
		logs.Critical("Oh, opps~")
	*/
	logs.Debug("This is Debug level log!")
	logs.Alert("This is Alert level log!")
	logs.Info("This is Inof level log!")
	logs.Warn("This is Warn level log!")
	logs.Error("This is Error level log")
	logs.Critical("This is Critical level log!")
	logs.Emergency("This is Emergency level log!")
	logs.Notice("This is Notice level log!")
	beego.Informational("This informational!")
	beego.BeeLogger.DelLogger("console")

}
