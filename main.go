package main

import (
	//"runtime"

	"github.com/gookit/color"
	"os"

	"github.com/EldersJavas/EbiBuilder/app"
	"github.com/EldersJavas/EbiBuilder/tool"
)

var App app.App

func main() {

	//////////////////////
	err := os.Chdir("examples")
	if err != nil {
		return
	}
	///////////////////////

	if color.SupportColor() != true {
		color.Disable()
	}
	tool.IsEbitenGame()
	App.Init()
	App.CmdApp.Run(os.Args[1:])
}
