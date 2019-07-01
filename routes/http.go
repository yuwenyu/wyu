package routes

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/yuwenyu/kernel"
	"wyu/app/http/controllers"
	"wyu/app/middleware"
)

type http struct {
	r    *gin.Engine
}

func New(r *gin.Engine) *http {
	return &http{r:r}
}

func (h *http) HttpRoutes() {
	h.r.Use(middleware.M())
	h.r.GET("/", controllers.NewIndexController().Ping)
	h.r.GET("/test.do", controllers.NewIndexController().Test)
}

func (h *http) HttpFuncMap() template.FuncMap {
	var i18nFunc kernel.I18N = kernel.NewI18n()
	return template.FuncMap{
		"T":i18nFunc.T,
		"C":middleware.ViewCfg,
	}
}
