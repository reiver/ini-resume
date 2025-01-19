package libtmpl

import (
	htmltemplate "html/template"
)

type EducationHTML struct {
	Credential string
	Name string
	Summary htmltemplate.HTML
	Topics string
	URL string
}
