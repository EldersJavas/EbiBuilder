package app

import (
	"fmt"
	"github.com/EldersJavas/EbiBuilder/tool"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/goutil/fsutil"
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
			c.StrOpt(&a, "buildmode", "m", "Debug", "the id option")
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
			v, err := tool.GetEbitenVer()
			if err != nil {
				tool.WarnPrint(err)
			}

			Pj.EbiVersion = v
			err = BuildGame(c)
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
	var IsFileBuild = false
	if Pj.EbiVersion == "" {
		tool.ErrorPrint("No Ebitengine version")
	} else if Pj.EbiVersion == "2.x" {
		IsFileBuild = true
	}
	tool.StepPrint("Build Start")
	switch Pj.BuildMode {

	///////////////////////////////
	// Debug
	case Debug:
		err := os.MkdirAll("output/Debug/", 777)
		if err != nil {
			return err
		}
		if IsFileBuild {
			_, err = sysutil.QuickExec(fmt.Sprintf("go build -o %v -tags=ebitendebug main.go", tool.ExecName("debug")))
		} else {
			_, err = sysutil.ExecCmd("go", []string{"build", "-tags=ebitendebug", "-o", tool.ExecName("debug")})
		}

		if err != nil {
			return err
		}

		err = fsutil.CopyFile(tool.ExecName("debug"), "output/Debug/"+tool.ExecName("debug"))
		if err != nil {
			return err
		}

		err = fsutil.DeleteIfExist(tool.ExecName("debug"))
		if err != nil {
			return err
		}

		tool.SuccessPrint("Debug build Success")

		////////////////////////////////

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
