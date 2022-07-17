package tool

import (
	"fmt"
	"github.com/gookit/color"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

func IsEbitenGame() bool {

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
			log.Println("No Ebitengine project in the dir.")
			return false
		}
	}
	for _, v := range strings.Split(string(f), "\n") {
		if strings.Contains(v, "github.com/hajimehoshi/ebiten/v2") {
			Versions = v[strings.LastIndex(v, "v2")+1:]
		}
	}
	if Versions == "" {
		return false
	}

	DebugPrint(fmt.Sprintf("ebiten version=%v", Versions))
	return true
}

func DebugPrint(s interface{}) {
	color.Debug.Printf("[Debug]%s: %v \n", time.Now().Format("2006/01/02 15:04:05.000"), s)
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

func ExecName(s string) string {
	switch runtime.GOOS {
	case `windows`:
		s = s + ".exe"
	}
	return s
}
