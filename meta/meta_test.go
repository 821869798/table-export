package meta

import (
	"bytes"
	"regexp"
	"testing"
	"text/template"
)

func TestMeta(t *testing.T) {
	reg := regexp.MustCompile(`^\[\](.+)$`)
	result := reg.FindAllStringSubmatch("[]int", -1)
	t.Log(result)
}

func TestTemplate(t *testing.T) {
	t1 := template.New("test1")
	tmpl, _ := t1.Parse(
		`
{{- define "T1"}}ONE {{println .}}{{end}}
{{- define "T2"}}{{template "T1" $}}{{end}}
{{- template "T2" . -}}
`)
	bb := &bytes.Buffer{}
	_ = tmpl.Execute(bb, "hello world")
	t.Log(bb.String())
}
