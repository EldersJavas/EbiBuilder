package app

import (
	"fmt"
	"github.com/EldersJavas/EbiBuilder/tool"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/sysutil"
	"os"
)

var Pj Project

var BuildCmd = NewBuildCmd()

func NewBuildCmd() *gcli.Command {
	a := &gcli.Command{
		Name:    "build",
		Desc:    "Build Ebitengine game",
		Aliases: []string{"Build", "BUILD", "buildgame"},
		Config: func(c *gcli.Command) {
			c.StrOpt(&Pj.BuildMode, "buildmode", "m", "Debug", "Build Mode")
			c.StrOpt(&Pj.OutputName, "outputname", "n", "Game", "Game output name")
			c.StrOpt(&Pj.Iconfile, "icon", "i", "favicon.ico", "icon file")
		},
		Examples: "ebibuilder build",
		Func: func(c *gcli.Command, args []string) error {
			v, err := tool.GetEbitenVer()
			if err != nil {
				tool.WarnPrint(err)
			}

			Pj.EbiVersion = v
			err = BuildGame()
			if err != nil {
				return err
			}
			return nil
		},
		Help: "Build the game",
	}
	return a
}

func BuildGame() error {
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
		if err := DebugBuild(IsFileBuild); err != nil {
			return err
		}
		tool.SuccessPrint("Debug build Success")
		// Release
	case Release:
		if err := ReleaseBuild(IsFileBuild); err != nil {
			return err
		}
		tool.SuccessPrint("Release build Success")

		///////////////////////////////////
	case All:
		err := DebugBuild(IsFileBuild)
		if err != nil {
			return err
		}
		tool.SuccessPrint("Debug build Success")
		err = ReleaseBuild(IsFileBuild)
		if err != nil {
			return err
		}
		tool.SuccessPrint("Release build Success")

	}
	return nil
}

func DebugBuild(IsFileBuild bool) error {
	tool.StepPrint("Build Debug Start")
	tool.StepPrint("Create dir...")
	err := os.MkdirAll("output/Debug/", 777)
	if err != nil {
		return err
	}
	tool.StepPrint("Building...")
	if IsFileBuild {
		_, err = sysutil.QuickExec(fmt.Sprintf("go build -o %v -tags=ebitendebug main.go", tool.ExecName(Pj.OutputName)))
	} else {
		_, err = sysutil.ExecCmd("go", []string{"build", "-tags=ebitendebug", "-o", tool.ExecName(Pj.OutputName)})
	}

	if err != nil {
		return err
	}
	tool.StepPrint("Managing...")
	err = fsutil.CopyFile(tool.ExecName(Pj.OutputName), "output/Debug/"+tool.ExecName(Pj.OutputName))
	if err != nil {
		return err
	}

	err = fsutil.DeleteIfExist(tool.ExecName(Pj.OutputName))
	if err != nil {
		return err
	}

	return nil
}

func ReleaseBuild(IsFileBuild bool) error {
	tool.StepPrint("Build Release Start")
	tool.StepPrint("Create dir...")
	err := os.MkdirAll("output/Release/", 777)
	if err != nil {
		return err
	}
	tool.StepPrint("Building...")
	if IsFileBuild {
		_, err = sysutil.QuickExec(fmt.Sprintf("go build -o %v main.go", tool.ExecName(Pj.OutputName)))
	} else {
		_, err = sysutil.ExecCmd("go", []string{"build", "-o", tool.ExecName(Pj.OutputName)})
	}

	if err != nil {
		return err
	}
	tool.StepPrint("Managing...")
	err = fsutil.CopyFile(tool.ExecName(Pj.OutputName), "output/Release/"+tool.ExecName(Pj.OutputName))
	if err != nil {
		return err
	}

	err = fsutil.DeleteIfExist(tool.ExecName(Pj.OutputName))
	if err != nil {
		return err
	}

	return nil
}
