package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arunraghunath/sharepics/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func processTemplate(w http.ResponseWriter, htmlfile string) {
	ts, err := views.Parse(htmlfile)
	if err != nil {
		log.Print(err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	ts.Execute(w, nil)
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
