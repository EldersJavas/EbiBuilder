// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"github.com/EldersJavas/EbiBuilder/tool"
	"github.com/gookit/gcli/v3"
)

var CreateCmd = NewCreateCmd()

func NewCreateCmd() *gcli.Command {
	a := &gcli.Command{
		Name:    "create",
		Desc:    "create new ebiengine project",
		Aliases: []string{"Create", "CREATE"},
		Config: func(c *gcli.Command) {

		},
		Examples: "ebibuilder create",
		Func: func(c *gcli.Command, args []string) error {
			v, err := tool.GetEbitenVer()
			if err != nil {
				tool.WarnPrint(err)
			}

			Pj.EbiVersion = v
			err = CreateGame()
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

	tool.SuccessPrint("Create Success")
	return nil
}
