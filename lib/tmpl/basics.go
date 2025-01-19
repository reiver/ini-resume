package libtmpl

import (
	htmltemplate "html/template"
)

type BasicsHTML struct {
	Name     string
	HeadLine htmltemplate.HTML
	Summary  htmltemplate.HTML
}
