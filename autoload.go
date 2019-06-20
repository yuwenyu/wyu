// Copyright 2019 YuwenYu.  All rights reserved.
// license that can be found in the LICENSE file.

package wyu

import (
	"strings"

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

		strResources := ad.kernel.Ini.K("template_root","resources").String()
		arrResources := strings.Split(strResources, "|")
		if arrResources == nil {
			panic("Templates Resources nil, Please check the configure!")
		}

		objTPL := ad.kernel.GinTemplate()
		for _, vResources := range arrResources {
			arrViews := strings.Split(vResources, ":")
			for _, view := range strings.Split(arrViews[1], ",") {
				arrTPL := ad.kernel.GinTemplateLoadByView(arrViews[0], view)
				objTPL.AddFromFilesFuncs(view, rHttp.HttpFuncMap(), arrTPL ...)
			}
		}
		r.HTMLRender = objTPL
	}

	strPort := ad.kernel.Ini.K("common_server","port").String()
	if strPort == "" {
		r.Run(":8080")
	} else {
		r.Run(strPort)
	}
}


