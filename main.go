package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func processTemplate(w http.ResponseWriter, htmlfile string) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	ts, err := template.ParseFiles(htmlfile)
	if err != nil {
		log.Print(err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	file := "./templates/home.html"
	processTemplate(w, file)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	file := "./templates/contact.html"
	processTemplate(w, file)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	file := "./templates/faq.html"
	processTemplate(w, file)
}

func urlParamHandler(w http.ResponseWriter, r *http.Request) {
	val := chi.URLParam(r, "param")
	if val == "" {
		fmt.Fprintf(w, "Empty param")
		return
	}
	fmt.Fprintf(w, "Param value is %s", val)
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", homeHandler)
	router.Get("/contact", contactHandler)
	router.Get("/faq", faqHandler)
	router.Get("/testurlparam", urlParamHandler)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", 404)
	})

	fmt.Println("Starting the server at :2020")
	http.ListenAndServe(":2020", router)
}
