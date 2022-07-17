package tool

import (
	"fmt"
	"github.com/EldersJavas/EbiBuilder/app"
	"github.com/gookit/color"
	"log"
	"os"
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

	DebugPrint(fmt.Sprintf("ebiten version=%v\n", Versions))
	return true
}

func DebugPrint(s string) {
	if app.DebugMode {
		color.Debug.Printf("[Debug]%v: %v", time.Now().Format("2006/01/02 15:04:05.000"), s)
	}
}

func InfoPrint(s string) {
	color.Info.Printf("[Info]%v: %v", time.Now().Format("2006/01/02 15:04:05.000"), s)
}
