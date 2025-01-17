package tmplsrv

import (
	htmltemplate "html/template"
	"io"

	"github.com/reiver/go-erorr"

	"github.com/reiver/ini-resume/lib/tmpl"
)

const (
	errNilTemplate = erorr.Error("nil template")
)

var htmltmpl *htmltemplate.Template = libtmpl.Default

func Render(writer io.Writer, src any) error {
	if nil == htmltmpl {
		return errNilTemplate
	}

	return htmltmpl.Execute(writer, src)
}
