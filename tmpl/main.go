package main

import (
	"bytes"
	"fmt"
	"text/template"
)

var tmplRaw = "{{range . }}{{.Title}}\n{{end}}"
var tmpl = template.Must(template.New("texttmpl").Parse(tmplRaw))

type Addr struct {
	Name string
	Title string
}

type NewAddr struct {
	Name string
	Title string
}

func main() {
	data := []*struct {
		Title string
		Name  string
		Value string
	}{
		{
			Title: "тайтл_1",
			Name:  "name_1",
			Value: "val_1",
		},
		{
			Title: "тайтл_2",
			Name:  "name_2",
			Value: "val_2",
		},
	}
	var out bytes.Buffer
	tmpl.Execute(&out, data)

	fmt.Println(out.String())

	addr := Addr{Name: "qwerty"}
	newaddr := NewAddr(addr)
	fmt.Printf("%+v", newaddr)
}
