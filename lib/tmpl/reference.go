package libtmpl

import (
	htmltemplate "html/template"
)

type ReferenceHTML struct {
	Quotation htmltemplate.HTML
	Citation  string
}
