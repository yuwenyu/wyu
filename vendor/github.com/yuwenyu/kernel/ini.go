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

/**
 * Todo: Initialize the ic.IniFile in the first place
 *
 * Method:Get
 * var ini INI = &ini{}
 * ini.Loading()
 * ini.LoadByFN(filename string)
 * ini.K(Section, Key).String()
 * ini.K(Section, Key).In("str", []string{"str", "str2"})
 * ini.K(Section, Key).MustInt(9999)
 * ini.K(Section, Key).MustBool(false)
**/

func (i *ini) K(section string, key string) *wyuIni.Key {
	if i.cfg == nil {
		panic("Error nil")
	}

	return i.cfg.Section(section).Key(key)
}

/**
 * Method:Set
 * ini.K(Section, Key).SetValue(value string).SaveTo(Path)
*/
//func (i *ini) SaveTo(fn string) *ini {
//	if i.cfg == nil {
//		panic("Error nil")
//	}
//
//	var Helper Helpers = &Helper{directory:i.directory,method:strIni}
//	i.cfg.SaveTo(Helper.TempCfgEnv(fn))
//	return i
//}

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