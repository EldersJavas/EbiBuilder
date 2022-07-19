package tool

import (
	"errors"
	"fmt"
	"github.com/gookit/color"
	"github.com/gookit/goutil/sysutil"
	"github.com/hashicorp/go-version"
	"os"
	"runtime"
	"strings"
	"time"
)

func CheckGo() bool {
	var Outgo, err = sysutil.QuickExec("go version")
	if err != nil {
		ErrorPrint("Unable to find Go")
		return false
	}
	te := strings.Split(Outgo, " ")
	Vergo, err := version.NewVersion(te[2][2:])
	if err != nil {
		ErrorPrint("Error go version:" + te[2][2:])
		return false
	}

	constraints, err := version.NewConstraint(">=1.15")
	if constraints.Check(Vergo) != true {
		return false
	}

	return true
}

// IsEbitenGame check ebitengine game project
func IsEbitenGame() (ebitengame bool, gomod bool) {

	var Versions string
	f, err := os.ReadFile("./go.mod")
	if err != nil {
		f2, err := os.ReadFile("./main.go")
		for _, v := range strings.Split(string(f2), "\n") {
			if strings.Contains(v, "github.com/hajimehoshi/ebiten/v2") {
				Versions = "2.x"
			}
		}
		if err != nil {
			WarnPrint("No Ebitengine project in the dir.")
			return false, false
		}
	}
	for _, v := range strings.Split(string(f), "\n") {
		if strings.Contains(v, "github.com/hajimehoshi/ebiten/v2") {
			Versions = v[strings.LastIndex(v, "v2")+1:]
		}
	}
	if Versions == "" {
		return false, false
	}
	//DebugPrint(fmt.Sprintf("ebiten version=%v", Versions))
	if Versions == "2.x" {
		return true, false
	}
	return true, true
}

// GetEbitenVer Get Ebitengine version
func GetEbitenVer() (string, error) {
	var Versions string
	f, err := os.ReadFile("./go.mod")
	if err != nil {
		f2, err := os.ReadFile("./main.go")
		for _, v := range strings.Split(string(f2), "\n") {
			if strings.Contains(v, "github.com/hajimehoshi/ebiten/v2") {
				Versions = "2.x"
			}
		}
		if err != nil {
			return "", errors.New("no Ebitengine project in the dir")
		}
	}
	for _, v := range strings.Split(string(f), "\n") {
		if strings.Contains(v, "github.com/hajimehoshi/ebiten/v2") {
			Versions = v[strings.LastIndex(v, "v2")+1:]
		}
	}
	if Versions == "" {
		return "", errors.New("unable to locate ebitengine version number")
	}

	DebugPrint(fmt.Sprintf("ebiten version=%v", Versions))
	return Versions, nil
}

func DebugPrint(s interface{}) {
	if DebugMode {
		color.Debug.Printf("[Debug]%s: %v \n", time.Now().Format("2006/01/02 15:04:05.000"), s)
	}
}

func InfoPrint(s interface{}) {
	color.Info.Printf("[Info]%s: %v \n", time.Now().Format("2006/01/02 15:04:05.000"), s)
}
func ErrorPrint(s interface{}) {
	color.Error.Printf("[Error]%s: %v \n", time.Now().Format("2006/01/02 15:04:05.000"), s)
}

func StepPrint(s interface{}) {
	color.FgMagenta.Printf("[Step]%s: %v \n", time.Now().Format("2006/01/02 15:04:05.000"), s)
}

func SuccessPrint(s interface{}) {
	color.Success.Printf("[Success]%s: %v \n", time.Now().Format("2006/01/02 15:04:05.000"), s)
}
func WarnPrint(s interface{}) {
	color.Warn.Printf("[Warn]%s: %v \n", time.Now().Format("2006/01/02 15:04:05.000"), s)
}

func ExecName(s string) string {
	switch runtime.GOOS {
	case `windows`:
		s = s + ".exe"
	}
	return s
}
