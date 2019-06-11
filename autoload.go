// Copyright 2020 YuwenYu.  All rights reserved.
// license that can be found in the LICENSE file.

package wyu

import (
	"github.com/yuwenyu/kernel"
	"wyu/routes"
)

func init() {
	ad := new()
	ad.running(":8081")
}

type Autoload struct {
	kernel *kernel.Kernel
}

func new() *Autoload {
	return &Autoload {
		kernel:kernel.New(),
	}
}

func (ad *Autoload) running(addr string) {
	//ad.static(ad.kernel.G)
	routes.New(ad.kernel.G).HttpRoutes()
	ad.kernel.Run(addr)
}


