package routes

import (
	"github.com/gin-gonic/gin"
	"wyu/app/http/controllers"
)

type http struct {
	r    *gin.Engine
}

func New(r *gin.Engine) *http {
	return &http{r:r}
}

func (h *http) HttpRoutes() {
	h.r.GET("/", controllers.NewIndexController().Ping)
	h.r.GET("/test.do", controllers.NewIndexController().Test)
}
