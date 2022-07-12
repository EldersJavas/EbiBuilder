package app

import "github.com/gookit/gcli/v3"

func NewBuildCmd() *gcli.Command {
	a := &gcli.Command{
		Flags:     gcli.Flags{},
		Arguments: gcli.Arguments{},
		Name:      "",
		Desc:      "",
		Aliases:   []string{},
		Category:  "",
		Config: func(c *gcli.Command) {
		},
		Hidden:   false,
		Subs:     []*gcli.Command{},
		Examples: "ebibuilder build",
		Func: func(c *gcli.Command, args []string) error {
			return nil
		},
		Help: "Build the game",
		HelpRender: func(c *gcli.Command) {
		},
	}
	return a
}
