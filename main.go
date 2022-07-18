package main

import (
	//"runtime"

	"github.com/EldersJavas/EbiBuilder/app"
	"github.com/EldersJavas/EbiBuilder/tool"
	"github.com/gookit/color"
)

var App = app.App

func main() {

	if tool.CheckGo() != true {
		tool.ErrorPrint("Go version less than v1.15")
	}
	if color.SupportColor() != true {
		color.Disable()
	}
	tool.IsEbitenGame()
	App.Run(nil)
}
