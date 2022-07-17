package app

import (
	"github.com/EldersJavas/EbiBuilder/tool"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/goutil/sysutil"
	"os"
)

var BuildCmd = NewBuildCmd()
var Pj Project

func NewBuildCmd() *gcli.Command {
	a := &gcli.Command{
		Flags:     gcli.Flags{},
		Arguments: gcli.Arguments{},
		Name:      "build",
		Desc:      "Build Ebitengine game",
		Aliases:   []string{"Build", "BUILD", "buildgame"},
		Category:  "",
		Config: func(c *gcli.Command) {
			a := ""
			c.StrOpt(&a, "buildmode", "m", "Release", "the id option")
			switch a {
			case "Debug":
				Pj.BuildMode = Debug
			case "Release":
				Pj.BuildMode = Release
			case "All":
				Pj.BuildMode = All

			}
		},
		Hidden:   false,
		Subs:     []*gcli.Command{},
		Examples: "ebibuilder build",
		Func: func(c *gcli.Command, args []string) error {
			err := BuildGame(c)
			if err != nil {
				return err
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

	out, err := sysutil.QuickExec("go version")
	if err != nil {
		tool.ErrorPrint(err)
		return err
	}
	tool.DebugPrint(out)
	tool.StepPrint("Build Debug Start")
	switch Pj.BuildMode {
	case Debug:
		err := os.MkdirAll("output/Debug/", 777)
		if err != nil {
			return err
		}
		//out, err := sysutil.QuickExec("go build -tags=ebitendebug")

		if err != nil {
			return err
		}
		tool.SuccessPrint("test")
	case Release:
		err := os.MkdirAll("output/Release/", 777)

		if err != nil {
			return err
		}
		tool.SuccessPrint("test")
	case All:
		err := os.MkdirAll("output/Debug/", 777)
		err = os.MkdirAll("output/Release/", 777)
		if err != nil {
			return err
		}
		tool.SuccessPrint("test")

	}
	return nil
}
