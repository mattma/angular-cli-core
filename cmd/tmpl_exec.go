package cmd

import (
	"bytes"
	"text/template"
)

type Parseble interface {
	Parse() string
}

//.gitignore parser

type GitIgnoreFile struct {
}

func (gi GitIgnoreFile) Parse() string {
	return parse_templates("gitignore", gitIgTmpl, gi)
}

//Main .go file parser

type MainFile struct {
	PackName string
}

func (m MainFile) Parse() string {
	return parse_templates("main", mainTmpl, m)
}

// Parse helper
func parse_templates(name, tmpl string, prs Parseble) string {
	w := &bytes.Buffer{}
	t := template.Must(template.New(name).Parse(tmpl))
	t.Execute(w, prs)
	return w.String()
}
