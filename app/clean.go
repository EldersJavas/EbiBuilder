// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"github.com/EldersJavas/EbiBuilder/tool"
	"github.com/gookit/gcli/v3"
	"os"
)

var CleanCmd = NewCleanCmd()

func NewCleanCmd() *gcli.Command {
	a := &gcli.Command{
		Name:    "clean",
		Desc:    "clean build",
		Aliases: []string{"Clean", "CLEAN"},
		Config: func(c *gcli.Command) {

		},
		Examples: "ebibuilder clean",
		Func: func(c *gcli.Command, args []string) error {
			v, err := tool.GetEbitenVer()
			if err != nil {
				tool.WarnPrint(err)
			}

			Pj.EbiVersion = v
			err = CleanGame(c)
			if err != nil {
				return err
			}
			return nil
		},
		Help: "clean build",
	}
	return a
}

func CleanGame(c *gcli.Command) error {
	err := os.RemoveAll("output")
	if err != nil {
		return err
	}
	return nil
}
