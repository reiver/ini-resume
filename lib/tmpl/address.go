package libtmpl

import (
	htmltemplate "html/template"
)

type AddressHTML struct {
	Label string
	Kind  string
	URI   htmltemplate.URL
}
