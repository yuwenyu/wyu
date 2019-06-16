package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type index struct {
	ctrl *Controller
}

func NewIndexController() *index {
	return &index {
		ctrl:NewController(),
	}
}

func (c *index) before() gin.HandlerFunc {
	return func(gc *gin.Context) {
		gc.Next()
		return
	}
}

func (c *index) Ping(gc *gin.Context) {
	gc.HTML(http.StatusOK, "index.html", gin.H{
		"title" : "test",
	})
}

func (c *index) Test(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{
		"msg": "test success",
		"message": http.Dir("my_file_system"),
		"params": gc.Request.Host,
	})
}
