package controllers

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	g *gin.Engine
}

func NewController() *Controller {
	return &Controller{}
}


