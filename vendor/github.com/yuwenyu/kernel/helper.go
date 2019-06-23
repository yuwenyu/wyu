/**
 * Copyright 2019 YuwenYu.  All rights reserved.
**/

package kernel

import (
	"os"
	"time"
)

const (
	StrCD string = "config" // directory of config (string)
	StrVirgule string = "/"
	StrUL string = "_"
	StrDOT string = "."
	StrColon string = ":"

	SysTimeFormat string = "2006-01-02 00:00:00"
	SysDateFormat string = "2006-01-02"
	DirDateFormat string = "20060102" // Directory Time Format
)

var (
	SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")
	MapTimeFormat map[string]string = map[string]string{
		"STF":time.Now().Format(SysTimeFormat),
		"SDF":time.Now().Format(SysDateFormat),
		"DDF":time.Now().Format(DirDateFormat),
	}
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

	strEnv := os.Getenv("WYU_ENV")
	if strEnv == "" {
		panic("Error ENV(WYU_ENV) Helper ...")
	}

	return h.directory + StrVirgule + h.method + StrVirgule + fn + StrDOT + strEnv + StrDOT + h.method
}