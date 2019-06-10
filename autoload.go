// Copyright 2020 YuwenYu.  All rights reserved.
// license that can be found in the LICENSE file.

package wyu

import (
	"fmt"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/yuwenyu/kernel"
	"wyu/routes"
)

func init() {
	ad := new()
	ad.running(":8081")
}

type Autoload struct {
	g *gin.Engine
	kernel *kernel.Kernel
}

func new() *Autoload {
	kernel := kernel.New()
	ad := &Autoload{
		kernel:kernel,
	}
	return ad
}

func (ad *Autoload) running(addr string) {
	//ad.kernel.SysLog()
	//ad.ini("config/")
	//ad.kernel.Ic.Dir = "config/"

	//ad.kernel.Ic.Loading()
	//fmt.Println(ad.kernel.Ic.SetIp("paths").K("data").String())
	//fmt.Println(ad.kernel.Ic.SetIp("paths").K("data").In("http", []string{"http", "https"}))

	r := gin.Default()
	r.HTMLRender = ad.tpl("resources/templates")
	ad.static(r)
	routes.New(r).HttpRoutes()
	r.Run(addr)
}

func (ad *Autoload) static(g *gin.Engine) {
	g.Static("/assets", "./resources/assets")
	g.StaticFile("/favicon.ico", "./resources/favicon.ico")
}

func (ad *Autoload) tpl(dir string) multitemplate.Renderer {
	tpl := multitemplate.NewRenderer()

	layout, err := filepath.Glob(dir + "/" + "layouts/wyu.html")
	if err != nil {
		panic(fmt.Sprintf("Template Layout-wyu Error: %s", err.Error()))
	}

	shareds, err := filepath.Glob(dir + "/" + "shared/*.html")
	if err != nil {
		panic(fmt.Sprintf("Template Shared-wyu Error: %s", err.Error()))
	}

	arrTPL := make([]string, 1)
	arrTPL  = append(layout, dir + "/views/index.html")

	for _, shared := range shareds {
		arrTPL = append(arrTPL, shared)
	}
	//log.Println(arrTPL)
	//views, err := filepath.Glob(dir + "/" + "views/*.html")
	//if err != nil {
	//	panic(fmt.Sprintf("Template view Error: %s", err.Error()))
	//}
	//
	//for _, view := range views {
	//	layoutCopy := make([]string, len(layouts))
	//	copy(layoutCopy, layouts)
	//	log.Println(layoutCopy)
	//	fs := append(layoutCopy, view)
	//	tpl.AddFromFiles(filepath.Base(view), fs ...)
	//	log.Println(fs)
	//}

	tpl.AddFromFiles("index.html", arrTPL ...)

	return tpl
}


