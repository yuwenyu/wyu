// Copyright 2019 YuwenYu.  All rights reserved.
// license that can be found in the LICENSE file.

package wyu

import (
	"github.com/yuwenyu/kernel"
	"wyu/routes"
)

func init() {
	new().running()
}

type autoload struct {
	kernel *kernel.Kernel
}

func new() *autoload {
	return &autoload {
		kernel:kernel.New(),
	}
}

func (ad *autoload) running() {
	r := ad.kernel.Run()

	/**
	 * TODO: Loading Templates
	 */
	bTpl, _ := ad.kernel.Ini.K("common_cfg","template_status").Bool()
	if bTpl {
		rHttp := routes.New(r)
		rHttp.HttpRoutes()

		objTPL := ad.kernel.GinTemplate()
		arrIndexTPL := ad.kernel.GinTemplateLoadByView("index")
		objTPL.AddFromFilesFuncs("index.html",rHttp.HttpFuncMap(), arrIndexTPL ...)
		arrTest1TPL := ad.kernel.GinTemplateLoadByView("test1")
		objTPL.AddFromFilesFuncs("test1.html",rHttp.HttpFuncMap(), arrTest1TPL ...)
		r.HTMLRender = objTPL

		//objTPL, arrTPL := ad.kernel.GinTemplate("index")
		//objTPL.AddFromFilesFuncs("index.html", rHttp.HttpFuncMap(), arrTPL ...)
		//objTPL, arrTPL := ad.kernel.GinTemplate("test1")
		//objTPL.AddFromFilesFuncs("test1.html", rHttp.HttpFuncMap(), arrTPL ...)
		//r.HTMLRender = objTPL
	}

	strPort := ad.kernel.Ini.K("common_server","port").String()
	if strPort == "" {
		r.Run(":8080")
	} else {
		r.Run(strPort)
	}
}


