// Copyright 2020 YuwenYu.  All rights reserved.
// license that can be found in the LICENSE file.

package wyu

import (
	"fmt"
	//"fmt"
	//"path/filepath"
	//
	//"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/yuwenyu/kernel"
	"github.com/yuwenyu/kernel/template"
	"wyu/routes"
)

func init() {
	ad := new()
	ad.running(":8081")
}

type Autoload struct {
	g *gin.Engine
	kernel *kernel.Kernel
	tpl *template.Template
}

func new() *Autoload {
	return &Autoload {
		kernel:kernel.New(),
		tpl:template.New(),
	}
}

func (ad *Autoload) running(addr string) {
	//ad.kernel.SysLog()

	ad.kernel.Ic.Loading()
	fmt.Println(ad.kernel.Ic.SetIp("template_root").K("directory").String())

	r := gin.Default()
	//r.HTMLRender = ad.tpl("resources/templates")
	r.HTMLRender = ad.tpl.SetDir("resources/templates").Tpl()
	ad.static(r)
	routes.New(r).HttpRoutes()
	r.Run(addr)
}

func (ad *Autoload) static(g *gin.Engine) {
	g.Static("/assets", "./resources/assets")
	g.StaticFile("/favicon.ico", "./resources/favicon.ico")
}


