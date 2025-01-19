package libtmpl

import (
	_ "embed"
	htmltemplate "html/template"
)

var Default *htmltemplate.Template = htmltemplate.Must(htmltemplate.New("resume").Parse(tmpl))

//go:embed tmpl.html
var tmpl string
