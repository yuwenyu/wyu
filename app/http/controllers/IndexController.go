package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type index struct {
	ctrl Controller
}

func NewIndexController() *index {
	return &index{}
}

func (i *index) before() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		return
	}
}

func (i *index) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s", time.Now().Format("2006_01_02")),
		"params": c.Params,
	})
}

func (i *index) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test success",
	})
}
