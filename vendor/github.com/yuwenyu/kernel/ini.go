/**
 * Copyright 2019 YuwenYu.  All rights reserved.
**/

package kernel

import (
	"path/filepath"

	wyuIni "gopkg.in/ini.v1"
)

const strIni string = "ini"

type INI interface {
	Loading() *ini
	LoadByFN(fn string) *ini
	K(section string, key string) *wyuIni.Key
}

type ini struct {
	directory string
	cfg *wyuIni.File
}

var _ INI = &ini{}

func (i *ini) initialize() {
	if i.directory == "" {
		i.directory = StrCD
	}
}

func (i *ini) Loading() *ini {
	i.initialize()

	var Helper Helpers = &Helper{directory:i.directory,method:strIni}
	fns, err := filepath.Glob(Helper.TempCfgEnv("*"))
	if err != nil {
		panic(err.Error())
	}

	if len(fns) == 0 {
		panic("No files is in the config directory.")
	}

	arrFns := make([]interface{}, len(fns))
	for k, fn := range fns {
		arrFns[k] = fn
	}

	cfg, err := wyuIni.Load(arrFns[0], arrFns ...)
	if err != nil {
		panic(err.Error())
	}

	i.cfg = cfg
	return i;
}

func (i *ini) LoadByFN(fn string) *ini {
	i.initialize()

	var Helper Helpers = &Helper{directory:i.directory,method:strIni}
	fns, err := filepath.Glob(Helper.TempCfgEnv(fn))
	if err != nil {
		panic(err.Error())
	}

	arrFns := make([]interface{}, len(fns))
	for k, fn := range fns {
		arrFns[k] = fn
	}

	cfg, err := wyuIni.Load(arrFns[0])
	if err != nil {
		panic(err.Error())
	}

	i.cfg = cfg
	return i;
}

/**
 * Todo: Initialize the ic.IniFile in the first place
 *
 * Method:Get
 * kernel.Ic.Loading()
 * kernel.Ic.K("Key")
 * kernel.Ic.K("Key").In("str", []string{"str1","str2"})
 * kernel.Ic.K("Key").MustInt(9999)
 * kernel.Ic.K("Key").MustBool(false)
 *
 * Method:Set
 * kernel.Ic.K("Key").SetValue("SetValue")
 * i.cfg.SaveTo(Path)
 * cfg.Section("").Key("app_mode").SetValue("production")
 * cfg.SaveTo(path+"my.ini.local")
 */
func (i *ini) K(section string, key string) *wyuIni.Key {
	if i.cfg == nil {
		panic("Error nil")
	}

	return i.cfg.Section(section).Key(key)
}

//func (i *Ic) SaveTo(key string, val string, fn string) *Ic {
//	if i.cfg == nil {
//		panic("Error nil")
//	}
//
//	i.K(key).SetValue(val)
//	i.cfg.SaveTo(i.h.TempCfgEnv(i.dir, sIni, fn))
//	return i
//}