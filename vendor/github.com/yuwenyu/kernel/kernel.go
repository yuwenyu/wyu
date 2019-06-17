/**
 * Copyright 2019 YuwenYu.  All rights reserved.
**/

package kernel

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {}

type Kernel struct {
	Ini *ini
}

func New() *Kernel {
	var ini INI = &ini{}
	return &Kernel{
		Ini:ini.Loading(),
	}
}

func (k *Kernel) Run() *gin.Engine {
	k.ginInitialize()

	r := gin.Default()
	r = k.ginTemplate(r)
	r = k.ginTemplateStatic(r)

	return r
}

func (k *Kernel) ginInitialize() {
	bLog, _ := k.Ini.K("common_cfg","log_status").Bool()
	if bLog {
		gin.DisableConsoleColor()

		cfgLogRoot := k.Ini.K("common_log","log_root").String()
		if cfgLogRoot == "" {
			cfgLogRoot = "storage" + StrVirgule + "logs" + StrVirgule
		}

		_, err := os.Stat(cfgLogRoot)
		if err != nil {
			panic(err.Error())
		}

		cfgLogPrefixFN := k.Ini.K("common_log","log_fn_prefix").String()
		if cfgLogPrefixFN == "" {
			cfgLogPrefixFN = "wyu"
		}

		fn := cfgLogRoot + StrVirgule + cfgLogPrefixFN + StrUL + MapTimeFormat["DDF"] + ".log"
		f, _ := os.Create(fn)
		gin.DefaultWriter = io.MultiWriter(f)
	} else {
		gin.ForceConsoleColor()
	}
}

func (k *Kernel) ginTemplate(r *gin.Engine) *gin.Engine {
	bTpl, _ := k.Ini.K("common_cfg","template_status").Bool()
	if bTpl {
		var templates templates = &template{
			directory:k.Ini.K("template_root","directory").String(),
		}
		r.HTMLRender = templates.Tpl()
	}

	return r
}

func (k *Kernel) ginTemplateStatic(r *gin.Engine) *gin.Engine {
	bTplStatic, _ := k.Ini.K("common_cfg", "template_static_status").Bool()
	if bTplStatic {
		static := k.Ini.K("template_statics","static").String()
		staticFile := k.Ini.K("template_statics","static_file").String()

		r.Static("/assets", static)
		r.StaticFile("/favicon.ico", staticFile)
	}

	return r
}
