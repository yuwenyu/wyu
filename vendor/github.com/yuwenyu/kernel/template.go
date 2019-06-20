/**
 * Copyright 2019 YuwenYu.  All rights reserved.
**/

package kernel

import (
	"fmt"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

type templates interface {
	Tpl() multitemplate.Renderer
	LoadingTPL(skeleton string, view string) []string
}

type template struct {
	directory string
}

var _ templates = &template{}

func (t *template) Tpl() multitemplate.Renderer {
	return multitemplate.NewRenderer()
}

func (t *template) LoadingTPL(skeleton string, view string) []string {
	if t.directory == "" {
		panic("Error: Empty Template Dir")
	}

	layout, err := filepath.Glob(t.directory + StrVirgule + "layouts" + StrVirgule + skeleton + ".html")
	if err != nil {
		panic(fmt.Sprintf("Template Layout-wyu Error: %s", err.Error()))
	}

	shareds, err := filepath.Glob(t.directory + StrVirgule + "shared" + StrVirgule + skeleton + StrVirgule + "*.html")
	if err != nil {
		panic(fmt.Sprintf("Template Shared-wyu Error: %s", err.Error()))
	}

	arrTPL := make([]string, 1)
	arrTPL  = append(layout, t.directory + StrVirgule + "views" + StrVirgule + skeleton + StrVirgule + view + ".html")

	for _, shared := range shareds {
		arrTPL = append(arrTPL, shared)
	}

	return arrTPL
}