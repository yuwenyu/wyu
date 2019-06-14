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

func (this *index) before() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		return
	}
}

func (this *index) Ping(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title" : "test",
	})
}

func (this *index) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "test success",
		"message": http.Dir("my_file_system"),
		"params": c.Params,
	})
}
