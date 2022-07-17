package main

import (
	//"runtime"

	"github.com/gookit/color"
	"os"

	"github.com/EldersJavas/EbiBuilder/app"
	"github.com/EldersJavas/EbiBuilder/tool"
)

var App = app.App

func main() {

	//////////////////////
	err := os.Chdir("examples")
	if err != nil {
		return
	}
	t, err := os.Getwd()
	if err != nil {
		return
	}
	tool.DebugPrint(t)

	///////////////////////

	if tool.CheckGo() != true {
		tool.ErrorPrint("Go version less than v1.15")
		panic("Go version less than v1.15")
	}
	if color.SupportColor() != true {
		color.Disable()
	}
	tool.IsEbitenGame()
	App.Run(nil)
}
