// Copyright 2020 YuwenYu.  All rights reserved.
// license that can be found in the LICENSE file.

package wyu

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/yuwenyu/kernel"
	"gopkg.in/ini.v1"

	"path/filepath"
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
	kl := kernel.New()
	return &Autoload{
		kernel:kl,
	}
}

func (ad *Autoload) running(addr string) {
	//ad.kernel.SysLog()
	//ad.ini("config/")
	//ad.kernel.Ic.Dir = "config/"
	ad.kernel.Ic.Loading()
	r := gin.Default()
	r.HTMLRender = ad.tpl("resources/templates")
	ad.static(r)
	routes.New(r).HttpRoutes()
	r.Run(addr)
}

func (ad *Autoload) ini(dir string) {

	//pathName, err := filepath.Glob(dir + "dev" + "/ini/my.ini")
	//if (err != nil) {
	//	panic("test error")
	//}



	pathName1, err := filepath.Glob(dir + "dev" + "/ini/*.ini")
	log.Println(pathName1)
	arr := make([]interface{}, len(pathName1))
	for v := range pathName1 {
		arr[v] = pathName1[v]
	}
	log.Println(arr)
	cfg, err := ini.Load(arr[0], arr ...)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	// 典型读取操作，默认分区可以使用空字符串表示
	log.Println(cfg.GetSection("paths"))
	fmt.Println("App Mode Test:", cfg.Section("").Key("app_mode").String())
	fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
	fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())

	// 我们可以做一些候选值限制的操作
	fmt.Println("Server Protocol:",
		cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))
	// 如果读取的值不在候选列表内，则会回退使用提供的默认值
	fmt.Println("Email Protocol:",
		cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"}))

	// 试一试自动类型转换
	fmt.Printf("Port Number: (%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(9999))
	fmt.Printf("Enforce Domain: (%[1]T) %[1]v\n", cfg.Section("server").Key("enforce_domain").MustBool(false))

	// 差不多了，修改某个值然后进行保存
	//cfg.Section("").Key("app_mode").SetValue("production")
	//cfg.SaveTo(path+"my.ini.local")
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


