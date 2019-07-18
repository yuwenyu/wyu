// Copyright 2019 YuwenYu.  All rights reserved.
// license that can be found in the LICENSE file.

package wyu

import (
	"io/ioutil"
	"strings"

	"github.com/yuwenyu/kernel"

	"wyu/config"
	"wyu/routes"
)

const (
	directoryView string = "resources/templates/views/"
	ginPort string = "8888"
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
	bTpl, _ := ad.kernel.Ini.K(
		config.MapConfLists[config.ConfCommons][0],
		config.MapConfParam[config.MapConfLists[config.ConfCommons][0]][0],
	).Bool()
	if bTpl {
		rHttp := routes.New(r)
		rHttp.HttpRoutes()

		strResources := ad.kernel.Ini.K(
			config.MapConfLists[config.ConfTemplates][1],
			config.MapConfParam[config.MapConfLists[config.ConfTemplates][1]][2],
		).String()
		if strResources == "" {
			panic("Templates Resources nil, Please check the configure!")
		}

		strDirectoryViews := ad.kernel.Ini.K(
			config.MapConfLists[config.ConfTemplates][1],
			config.MapConfParam[config.MapConfLists[config.ConfTemplates][1]][1],
		).String()
		if strDirectoryViews == "" {
			strDirectoryViews = directoryView
		}

		arrResources := strings.Split(strResources, kernel.StrColon)
		objTPL := ad.kernel.GinTemplate()
		for _, skeleton := range arrResources {
			views, _ := ioutil.ReadDir(strDirectoryViews + skeleton)
			for _, view := range views {
				arrTPL := ad.kernel.GinTemplateLoadByView(skeleton, view.Name())
				objTPL.AddFromFilesFuncs(view.Name(), rHttp.HttpFuncMap(), arrTPL ...)
			}
		}
		r.HTMLRender = objTPL
	}

	strPort := ad.kernel.Ini.K(
		config.MapConfLists[config.ConfCommons][2],
		config.MapConfParam[config.MapConfLists[config.ConfCommons][2]][0],
	).String()
	if strPort == "" {
		r.Run(kernel.StrColon + ginPort)
	} else {
		r.Run(kernel.StrColon + strPort)
	}
}