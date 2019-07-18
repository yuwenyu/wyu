package middleware

import (
	"github.com/yuwenyu/kernel"
	"wyu/config"
)

type View struct {}

func ViewCfg(key string) string {
	var c kernel.INI = kernel.NewIni().LoadByFN(config.ConfTemplates)
	return c.K(config.MapConfLists[config.ConfTemplates][2], key).String()
}