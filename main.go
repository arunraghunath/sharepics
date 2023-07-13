package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to Sharepics website!!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>You may contact me at <a href=\"mailto:test@gmail.com\">myemail</a></p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	w.Write([]byte("<h1>FAQ Page</h1><p>All help you need can be found here</p>"))
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
