// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"embed"
	"fmt"
	"github.com/EldersJavas/EbiBuilder/tool"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/sysutil"
	"github.com/valyala/fasttemplate"
	"io"
	"io/fs"
	"os"
)

var CreateCmd = NewCreateCmd()

//go:embed demo
var CreateFile embed.FS
var CFF fs.FS
var Npj Project

func init() {
	f, err := fs.Sub(CreateFile, "demo")
	if err != nil {
		panic(err)
	}
	CFF = f
}
func NewCreateCmd() *gcli.Command {
	a := &gcli.Command{
		Name:    "create",
		Desc:    "create new ebiengine project",
		Aliases: []string{"Create", "CREATE"},
		Config: func(c *gcli.Command) {
			c.StrOpt(&Npj.GameName, "gamename", "n", "Game", "Game project name")
			c.BoolOpt(&Npj.IsGomod, "gomod", "gm", true, "Go mod")
			//c.StrOpt(&Npj.Iconfile, "icon", "i", "favicon.ico", "icon file")
		},
		Examples: "ebibuilder create",
		Func: func(c *gcli.Command, args []string) error {
			err := CreateGame()
			if err != nil {
				return err
			}
			return nil
		},
		Help: "create a new ebiengine project.",
	}
	return a
}

func CreateGame() error {
	tool.StepPrint("Create Start")
	err := fsutil.Mkdir(Npj.GameName, 0777)
	if err != nil {
		return err
	}
	err = os.Chdir(Npj.GameName)
	if err != nil {
		return err
	}
	fi, err := CFF.Open("main.go.tmpl")

	if err != nil {
		return err
	}
	asll, err := io.ReadAll(fi)
	if err != nil {
		return err
	}
	t := fasttemplate.New(string(asll), "{{", "}}")
	s := t.ExecuteString(map[string]interface{}{
		"GameName": Npj.GameName,
	})
	/*tmpl := template.New("demo")
	tmpl, err = tmpl.Parse(string(asll))
	if err != nil {
		return err
	}
	err = tmpl.Execute(Wr, p)
	if err != nil {
		return err
	}*/
	err = os.WriteFile("main.go", []byte(fmt.Sprint(s)), 0777)
	if err != nil {
		return err
	}
	pp, err := sysutil.QuickExec("go mod init " + Npj.GameName)
	tool.InfoPrint(pp)
	if err != nil {
		return err
	}
	pp, err = sysutil.QuickExec("go mod tidy")
	tool.InfoPrint(pp)
	if err != nil {
		return err
	}
	err = InitConfig()
	if err != nil {
		return err
	}

	tool.SuccessPrint("Create Success")
	return nil
}
