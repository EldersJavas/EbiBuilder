// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"github.com/EldersJavas/EbiBuilder/tool"
	"github.com/gookit/color"
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

			c.AddSubs(&gcli.Command{
				Name:     "print",
				Desc:     "print project",
				Aliases:  []string{"Print", "PRINT"},
				Config:   nil,
				Examples: "ebibuilder config print",
				Func: func(c *gcli.Command, args []string) error {
					err := PrintConfig()
					if err != nil {
						return err
					}
					return nil
				},
				Help: "print project",
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

func PrintConfig() error {
	tool.StepPrint("Print config Start")
	file, err := os.ReadFile("build.json")
	if err != nil {
		return err
	}
	proj, err := UnmarshalProject(file)
	if err != nil {
		return err
	}
	bm := "Debug"
	switch proj.BuildMode {
	case Debug:
		bm = "Debug"
	case Release:
		bm = "Debug"
	case All:
		bm = "Debug,Release"
	}
	color.Printf(
		`

<cyan>Go mod</>:     		 <red>%v</>
<cyan>Ebiten v1</>:  		 <red>%v</>
<cyan>Ebitengine version</>:      <yellow>%v</> 
<cyan>Build Mode</>:   		 %v
<cyan>Output Name</>:  		 %v
<cyan>Config</>:      		 %v
<cyan>Path</>:        		 %v
<cyan>Game Name</>:    		 %v
<cyan>Game version</>: 		 <yellow>%v</>

`, proj.IsGomod, proj.IsEbitenv1, proj.EbiVersion, bm, proj.OutputName, proj.Config, proj.Path, proj.GameName, proj.GameVersion)
	tool.SuccessPrint("Print config Success")
	return nil
}
