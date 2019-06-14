package kernel

import (
	"fmt"
	"path/filepath"

	pkgIni"gopkg.in/ini.v1"
)

const strIni string = "ini"

type INI interface {
	Loading() *ini
	K(section string, key string) *pkgIni.Key
}

type ini struct {
	directory string
	cfg *pkgIni.File
}

var _ INI = &ini{}

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
func (i *ini) K(section string, key string) *pkgIni.Key {
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

func (i *ini) Loading() *ini {
	var Helper Helpers = &Helper{directory:i.directory,method:strIni}

	fns, err := filepath.Glob(Helper.TempCfgEnv("*"))
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err.Error()))
	}

	if len(fns) == 0 {
		panic(fmt.Sprintf("Error: %s", err.Error()))
	}

	arrFns := make([]interface{}, len(fns))
	for k, fn := range fns {
		arrFns[k] = fn
	}

	arrFns = append(arrFns, i.directory + "templates" + "." + strIni)
	arrFns = append(arrFns, i.directory + "commons" + "." + strIni)

	cfg, err := pkgIni.Load(arrFns[0], arrFns ...)
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err.Error()))
	}

	i.cfg = cfg

	return i;
}