// Copyright 2020 YuwenYu.  All rights reserved.
// license that can be found in the LICENSE file.

package wyu

import (
	"github.com/gin-gonic/gin"
	"github.com/yuwenyu/kernel"
	"wyu/routes"
)

func init() {
	new().running(":8081")
}

type Autoload struct {
	kernel kernel.Kernel
}

func new() *Autoload {
	return &Autoload{}
}

func (ad *Autoload) running(addr string) {
	kernel.New().SysLog()

	r := gin.Default()
	routes.New(r).HttpRoutes()
	r.Run(addr)
}
