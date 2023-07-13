package controllers

import (
	"html/template"
	"net/http"

	"github.com/arunraghunath/sharepics/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{Question: "q1",
			Answer: "a1"},
		{Question: "q2",
			Answer: "a2"},
		{Question: "q3",
			Answer: "a3"},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
