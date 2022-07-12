package main

import (
	//"runtime"
	"github.com/EldersJavas/EbiBuilder/app"
)

var App app.App

// 测试运行: go run ./_examples/cliapp.go && ./cliapp
func main() {
	App.Init()
	App.CmdApp.Run(nil)
}
