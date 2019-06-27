package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wyu/app/http/services"
)

type index struct {
	ctrl *Controller
	s services.IndexService
}

func NewIndexController() *index {
	return &index {
		ctrl:NewController(),
		s:services.NewIndexService(),
	}
}

func (c *index) Ping(gc *gin.Context) {
	gc.HTML(http.StatusOK, "index.html", gin.H{
		"title" : "test",
		"redis" : c.s.GetRedis(),
	})
}

func (c *index) Test(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{
		"msg": "test success",
		"message": http.Dir("my_file_system"),
		"params": gc.Request.Host,
		"datasource": c.s.FetchAll(),
		"redis": c.s.GetRedis(),
	})
}
