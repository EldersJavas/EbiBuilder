package app

import (
	"github.com/gookit/gcli/v3"
	"github.com/gookit/gcli/v3/builtin"
)

type App struct{
	CmdApp *gcli.App
	Path string
}



func (a *App)Init(){
	a.CmdApp = gcli.NewApp(func(app *gcli.App) {
		app.Version = VERSION
		app.Desc = "this is my cli application"
		app.On(gcli.EvtAppInit, func(data ...interface{}) bool {
			// do something...
			// fmt.Println("init app")
			return false
		})
	
		// app.SetVerbose(gcli.VerbDebug)
		// app.DefaultCommand("example")
		app.Logo.Text = `
		
███████╗██████╗ ██╗██████╗ ██╗   ██╗██╗██╗     ██████╗ ███████╗██████╗ 
██╔════╝██╔══██╗██║██╔══██╗██║   ██║██║██║     ██╔══██╗██╔════╝██╔══██╗
█████╗  ██████╔╝██║██████╔╝██║   ██║██║██║     ██║  ██║█████╗  ██████╔╝
██╔══╝  ██╔══██╗██║██╔══██╗██║   ██║██║██║     ██║  ██║██╔══╝  ██╔══██╗
███████╗██████╔╝██║██████╔╝╚██████╔╝██║███████╗██████╔╝███████╗██║  ██║
╚══════╝╚═════╝ ╚═╝╚═════╝  ╚═════╝ ╚═╝╚══════╝╚═════╝ ╚══════╝╚═╝  ╚═╝
                                                                       
		`
	})
	//a.CmdApp.Add()
	// a.CmdApp.Add(&gcli.Command{
	// 	Name: "demo",
	// 	// allow color tag and {$cmd} will be replace to 'demo'
	// 	Desc: "this is a description <info>message</> for command", 
	// 	Aliases: []string{"dm"},
	// 	Func: func (cmd *gcli.Command, args []string) error {
	// 		gcli.Println("hello, in the demo command")
	// 		return nil
	// 	},
	// })
	a.CmdApp.Add(builtin.GenAutoComplete())
	// .... add more ...
	}
