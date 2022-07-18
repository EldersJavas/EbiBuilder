// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"github.com/EldersJavas/EbiBuilder/tool"
	"github.com/blang/semver"
	"github.com/gookit/gcli/v3"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

var SelfUpdateCmd = NewSelfUpdateCmd()

func NewSelfUpdateCmd() *gcli.Command {
	a := &gcli.Command{
		Name:    "self-update",
		Desc:    "self-update",
		Aliases: []string{"Self-update", "selfupdate", "Selfupdate"},
		Config: func(c *gcli.Command) {

		},
		Examples: "ebibuilder self-update",
		Func: func(c *gcli.Command, args []string) error {
			v, err := tool.GetEbitenVer()
			if err != nil {
				tool.WarnPrint(err)
			}

			Pj.EbiVersion = v
			err = SelfUpdate()
			if err != nil {
				return err
			}
			return nil
		},
		Help: "self-update",
	}
	return a
}

func SelfUpdate() error {
	version := VERSION
	slug := "EldersJavas/EbiBuilder"
	tool.StepPrint("Self-update Start")
	selfupdate.EnableLog()

	previous := semver.MustParse(version)
	latest, err := selfupdate.UpdateSelf(previous, slug)
	if err != nil {
		return err
	}
	if previous.Equals(latest.Version) {
		tool.InfoPrint("Current binary is the latest version")
	} else {
		tool.SuccessPrint("Update successfully done to version" + latest.Version.String())
		tool.InfoPrint("Release note:\n" + latest.ReleaseNotes)
	}
	return nil
}
