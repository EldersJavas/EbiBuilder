package main

import (
	//"runtime"

	"github.com/gookit/color"

	"github.com/EldersJavas/EbiBuilder/app"
	"github.com/EldersJavas/EbiBuilder/tool"
)

var App app.App

// 测试运行: go run ./_examples/cliapp.go && ./cliapp
func main() {
	if color.SupportColor() {
		color.Disable()
	}
	tool.IsEbitenGame()
	App.Init()
	App.CmdApp.Run(nil)
}
