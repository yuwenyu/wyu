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
	routes.New(r).HttpRoutes()

	strPort := ad.kernel.Ini.K("common_server","port").String()
	if strPort == "" {
		r.Run(":8080")
	} else {
		r.Run(strPort)
	}
}


