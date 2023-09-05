package controllers

import (
	"html/template"
	"net/http"

	"github.com/vasfvitor/WebGo/views"
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
		{
			Question: "Is there a free version?",
			Answer:   "Yes",
		},
		{
			Question: "Where?",
			Answer:   "Here",
		},
		{
			Question: "Can you help me?",
			Answer:   "Yes",
		},
		{
			Question: "HTML?",
			Answer:   `<a href="/">Yep</a>`,
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
