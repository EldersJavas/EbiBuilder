package main

import (
	//"runtime"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/gcli/v3/_examples/cmd"
	"github.com/gookit/gcli/v3/builtin"
)

// 测试运行: go run ./_examples/cliapp.go && ./cliapp
func main() {
    app := gcli.NewApp(func(app *gcli.App) {
		app.Version = "3.0.0"
		app.Desc = "this is my cli application"
		app.On(gcli.EvtAppInit, func(data ...interface{}) bool {
			// do something...
			// fmt.Println("init app")
			return false
		})

		// app.SetVerbose(gcli.VerbDebug)
		// app.DefaultCommand("example")
		app.Logo.Text = ``
	})
    app.Add(cmd.Example)
    app.Add(&gcli.Command{
        Name: "demo",
        // allow color tag and {$cmd} will be replace to 'demo'
        Desc: "this is a description <info>message</> for command", 
        Aliases: []string{"dm"},
        Func: func (cmd *gcli.Command, args []string) error {
            gcli.Println("hello, in the demo command")
            return nil
        },
    })
	app.Add(builtin.GenAutoComplete())
    // .... add more ...

    app.Run(nil)
}