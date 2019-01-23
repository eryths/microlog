package component

import (
	"html/template"
	"io"
)

type Layout struct {
	tmpl *template.Template
}

func CreateView(layout string, tpl string) Layout {
	return Layout{
		tmpl: template.Must(template.ParseFiles("web/tpl/layout.html", "web/tpl/users.html")),
	}
}

func (v *Layout) Render(wr io.Writer) error {
	return v.tmpl.Execute(wr, nil)
}
