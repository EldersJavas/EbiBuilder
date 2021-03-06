package app

import (
	"github.com/gookit/gcli/v3"
	"github.com/gookit/gcli/v3/builtin"
)

var App = gcli.NewApp(func(app *gcli.App) {
	app.Version = VERSION
	app.Desc = "A tool for managing & building the ebitengine game."
	app.On(gcli.EvtAppInit, func(data ...interface{}) bool {
		return false
	})
	app.Logo.Text = `
		
███████╗██████╗ ██╗██████╗ ██╗   ██╗██╗██╗     ██████╗ ███████╗██████╗ 
██╔════╝██╔══██╗██║██╔══██╗██║   ██║██║██║     ██╔══██╗██╔════╝██╔══██╗
█████╗  ██████╔╝██║██████╔╝██║   ██║██║██║     ██║  ██║█████╗  ██████╔╝
██╔══╝  ██╔══██╗██║██╔══██╗██║   ██║██║██║     ██║  ██║██╔══╝  ██╔══██╗
███████╗██████╔╝██║██████╔╝╚██████╔╝██║███████╗██████╔╝███████╗██║  ██║
╚══════╝╚═════╝ ╚═╝╚═════╝  ╚═════╝ ╚═╝╚══════╝╚═════╝ ╚══════╝╚═╝  ╚═╝
                                                                       
		`
	app.Add(SelfUpdateCmd, CleanCmd, BuildCmd, ConfigCmd, CreateCmd)
	app.Add(builtin.GenAutoComplete())
	app.Add(builtin.GenEmojiMap)
})
