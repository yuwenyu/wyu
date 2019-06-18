package kernel

import (
	"github.com/syyongx/ii18n"
)

type I18N interface {
	T(cg string, key string, ln string) string
}

type i18n struct {
	cfg cfgII18N
}

var _ I18N = &i18n{}

type cfgII18N map[string]ii18n.Config

func NewI18n() *i18n {
	return &i18n{}
}

func (thisI18n *i18n) initialize() *i18n {
	thisI18n.cfg = cfgII18N{
		"app":{
			SourceNewFunc: ii18n.NewJSONSource,
			OriginalLang:  "en",
			BasePath:      "resources/lang",
			FileMap: map[string]string{
				"app":   "app.json",
			},
		},
	}
	ii18n.NewI18N(thisI18n.cfg)
	return thisI18n
}

func (thisI18n *i18n) T(cg string, key string, ln string) string {
	thisI18n.initialize()
	return ii18n.T(cg, key, nil, ln)
}


