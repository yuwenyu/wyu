package kernel

import (
	"os"
)

const (
	strCfgDirectory string = "config"
	strVirgule string = "/"
	strRTN string = "\n"
)

type Helpers interface {
	TempCfgEnv(fn string) string
}

type Helper struct {
	directory string
	method string
}

var _ Helpers = &Helper{}

func (h *Helper) TempCfgEnv(fn string) string {
	if h.directory == "" || h.method == "" {
		panic("Error Empty Helper ...")
	}

	return h.directory + strVirgule + h.method + strVirgule + os.Getenv("WYU_ENV") + strVirgule + fn + "." + h.method
}