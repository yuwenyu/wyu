package kernel

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {}

type Kernel struct {
	Ini *ini
}

func New() *Kernel {
	var ini INI = &ini{directory:strCfgDirectory + strVirgule}
	return &Kernel{
		Ini:ini.Loading(),
	}
}

func (k *Kernel) Run() *gin.Engine {
	bLog, _ := k.Ini.K("common_cfg","log_status").Bool()
	if bLog {
		gin.DisableConsoleColor()

		fn := "storage/logs/wyu_" + time.Now().Format("2006_01_01") + ".log"
		f, _ := os.Create(fn)
		gin.DefaultWriter = io.MultiWriter(f)
	}

	r := gin.Default()

	bTpl, _ := k.Ini.K("common_cfg","template_status").Bool()
	if bTpl {
		//var templates templates = &template{
		//	directory:k.Ic.SetIp("template_root").K("directory").String(),
		//}
		var templates templates = &template{
			directory:k.Ini.K("template_root","directory").String(),
		}
		r.HTMLRender = templates.Tpl()

		//static := k.Ic.SetIp("template_statics").K("static").String()
		//staticFile := k.Ic.SetIp("template_statics").K("static_file").String()
		//
		//r.Static("/assets", static)
		//r.StaticFile("/favicon.ico", staticFile)
	}

	return r
}
