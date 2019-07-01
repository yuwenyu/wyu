package middleware

import (
	"github.com/yuwenyu/kernel"
)

type View struct {}

func ViewCfg(key string) string {
	var c kernel.INI = kernel.NewIni()
	c.LoadByFN("templates")
	return c.K("template_config", key).String()
}