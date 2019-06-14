package controllers

import (
	"github.com/yuwenyu/kernel"
)

type Controller struct {
	K *kernel.Kernel
}

func NewController() *Controller {
	return &Controller{
		kernel.New(),
	}
}


