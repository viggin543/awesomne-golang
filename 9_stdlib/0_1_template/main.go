package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"
)

//go:embed foo.suffix.gohtml
var templateString string

func main() {
	t := template.Must(template.New("dontcare").Parse(templateString))
	var tpl bytes.Buffer
	err := t.Execute(&tpl, struct {
		FOO string
		BAR []string
		BAZ map[string]string
	}{
		FOO: "FUN",
		BAR: []string{"THIS", "IS", "SO", "COOL"},
		BAZ: map[string]string{
			"THIS": "1",
			"IS":   "2",
			"SO":   "3",
			"COOL": "4",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(tpl.String())
}
