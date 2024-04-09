package generator

import (
	_ "embed"

	"fmt"

	"io"
	
	"text/template"

	"github.com/jsightapi/jsight-api-core/kit"
)

//go:embed template.html
var templateContent string

type html struct {
	common
}

var _ Generator = html{}

func newHTML() html {
	h := html{}
	h.gen = h.generate
	return h
}

func (html) generate(japi kit.JApi, out io.Writer) error {
	json, err := japi.ToJson()
	if err != nil {
		return fmt.Errorf("generate JSON: %w", err)
	}
	tmpl, err := template.New("").Parse(templateContent)
	if err != nil {
		return fmt.Errorf("parse template: %w", err)
	}

	return tmpl.Execute(out, map[string]string{
		"jdoc": string(json),
	})
}
