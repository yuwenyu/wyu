// Copyright 2020 YuwenYu.  All rights reserved.
// license that can be found in the LICENSE file.

package wyu

import (
	"github.com/yuwenyu/kernel"
	"wyu/routes"
)

func init() {
	new().running(":8081")
}

type autoload struct {
	kernel *kernel.Kernel
}

func new() *autoload {
	return &autoload {
		kernel:kernel.New(),
	}
}

func (ad *autoload) running(addr string) {
	r := ad.kernel.Run()
	routes.New(r).HttpRoutes()
	r.Run(addr)
}

//func (ad *Autoload) T() {
//	fmt.Println(ad.kernel.I18n.T("test","cn"))
//	//config := map[string]ii18n.Config{
//	//	"app": {
//	//		SourceNewFunc: ii18n.NewJSONSource,
//	//		OriginalLang:  "en",
//	//		BasePath:      "resources/lang",
//	//		FileMap: map[string]string{
//	//			"app":   "app.json",
//	//		},
//	//	},
//	//}
//	//
//	//ii18n.NewI18N(config)
//	//test := ii18n.T("app", "test", nil, "cn")
//	//fmt.Println(test)
//}


