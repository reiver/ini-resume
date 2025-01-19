package libtmpl

import (
	htmltemplate "html/template"
)

type WorkHTML struct {
	Highlights []htmltemplate.HTML
	OrgHeadLine  htmltemplate.HTML
	OrgName      htmltemplate.HTML
	OrgSummary   htmltemplate.HTML
	Skills     []string
	Summary      htmltemplate.HTML
	Title        htmltemplate.HTML
	WhenEnded    string
	WhenStarted  string
}
