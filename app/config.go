// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"github.com/EldersJavas/EbiBuilder/tool"
	"github.com/gookit/gcli/v3"
	"os"
)

var ConfigCmd = NewConfigCmd()

func NewConfigCmd() *gcli.Command {
	a := &gcli.Command{
		Name:    "config",
		Desc:    "config",
		Aliases: []string{"Config", "CONFIG"},
		Config: func(c *gcli.Command) {
			// init config
			c.AddSubs(&gcli.Command{
				Name:     "init",
				Desc:     "init config",
				Aliases:  []string{"Init", "INIT"},
				Config:   nil,
				Examples: "ebibuilder config init",
				Func: func(c *gcli.Command, args []string) error {
					err := InitConfig()
					if err != nil {
						return err
					}
					return nil
				},
				Help: "init config",
			})
		},
		Examples: "ebibuilder config",
		Func: func(c *gcli.Command, args []string) error {
			v, err := tool.GetEbitenVer()
			if err != nil {
				tool.WarnPrint(err)
			}
			Pj.EbiVersion = v
			err = ConfigGame()
			if err != nil {
				return err
			}
			return nil
		},

		Help: "config ebitengine build project",
	}
	return a
}

func ConfigGame() error {
	return nil
}

func InitConfig() error {
	tool.StepPrint("Init config Start")
	a, b := tool.IsEbitenGame()
	v, err := tool.GetEbitenVer()
	if err != nil {
		return err
	}
	if a != true {
		tool.WarnPrint("Not ebitengine game")
	}
	p := Project{
		IsGomod:     b,
		IsEbitenv1:  false,
		EbiVersion:  v,
		BuildMode:   Debug,
		OutputName:  "output",
		Config:      "",
		Path:        "",
		GameName:    "game",
		GameVersion: "1.0",
	}
	j, err := p.Marshal()
	if err != nil {
		return err
	}
	err = os.WriteFile("build.json", j, 777)
	if err != nil {
		return err
	}
	tool.SuccessPrint("Init Config Success, file name:config.json")
	return nil
}
