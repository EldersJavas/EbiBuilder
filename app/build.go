package app

import (
	"github.com/gookit/gcli/v3"
	"github.com/gookit/goutil/sysutil"
)

var BuildCmd = NewBuildCmd()
var Pj Project

func NewBuildCmd() *gcli.Command {
	a := &gcli.Command{
		Flags:     gcli.Flags{},
		Arguments: gcli.Arguments{},
		Name:      "build",
		Desc:      "Build Ebitengine game",
		Aliases:   []string{},
		Category:  "",
		Config: func(c *gcli.Command) {
			a := ""
			c.StrOpt(&a, "buildmode", "m", "Release", "the id option")
			switch a {
			case "Debug":
				Pj.BuildMode = Debug
			case "Release":
				Pj.BuildMode = Release
			}
		},
		Hidden:   false,
		Subs:     []*gcli.Command{},
		Examples: "ebibuilder build",
		Func: func(c *gcli.Command, args []string) error {
			if Pj.BuildMode == 0 {
				err := BuildGame(c)
				if err != nil {
					return err
				}
			}
			return nil
		},
		Help: "Build the game",
		HelpRender: func(c *gcli.Command) {
		},
	}
	return a
}

func BuildGame(c *gcli.Command) error {

	out, err := sysutil.QuickExec("git")

	if err != nil {
		return err
	}
	c.Infoln(out)
	return nil

}
