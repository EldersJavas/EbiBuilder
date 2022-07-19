// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"fmt"
	"github.com/EldersJavas/EbiBuilder/tool"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/sysutil"
	"os"
	"syscall"
)

// TODO: Publish cmd

func PublishWASM() error {
	var IsFileBuild = false
	if Pj.EbiVersion == "" {
		tool.ErrorPrint("No Ebitengine version")
	} else if Pj.EbiVersion == "2.x" {
		IsFileBuild = true
	}
	tool.StepPrint("Publish Start")

	err := syscall.Setenv("GOOS", "js")
	if err != nil {
		return err
	}
	err = syscall.Setenv("GOARCH", "wasm")
	if err != nil {
		return err
	}

	err = os.MkdirAll("publish/wasm", 777)
	if err != nil {
		return err
	}

	if IsFileBuild {
		_, err = sysutil.QuickExec(fmt.Sprintf("go build -o %v main.go", Pj.OutputName))
	} else {
		_, err = sysutil.ExecCmd("go", []string{"build", "-o", Pj.OutputName})
	}
	if err != nil {
		return err
	}
	err = fsutil.CopyFile(Pj.OutputName+".wasm", "publish/wasm/"+Pj.OutputName+".wasm")
	if err != nil {
		return err
	}

	err = fsutil.DeleteIfExist(Pj.OutputName + ".wasm")
	if err != nil {
		return err
	}

	tool.SuccessPrint("Publish wasm Success")
	return nil
}
