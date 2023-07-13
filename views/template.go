package views

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func Parse(file string) (t Template, err error) {
	htmlTpl, err := template.ParseFiles(file)
	if err != nil {
		return Template{}, fmt.Errorf("Parsing template %w", err)
	}
	return Template{
		htmlTpl: htmlTpl,
	}, nil

}

func ParseFS(fs embed.FS, file ...string) (t Template, err error) {
	htmlTpl, err := template.ParseFS(fs, file...)
	if err != nil {
		return Template{}, fmt.Errorf("Parsing template %w", err)
	}
	return Template{
		htmlTpl: htmlTpl,
	}, nil
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("Executing template %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}
