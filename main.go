package main

import (
	//"runtime"

	"github.com/gookit/color"

	"github.com/EldersJavas/EbiBuilder/app"
	"github.com/EldersJavas/EbiBuilder/tool"
)

var App app.App

func main() {
	if color.SupportColor() != true {
		color.Disable()
	}
	tool.IsEbitenGame()
	App.Init()
	App.CmdApp.Run(nil)
}
