package main

import (
	"fmt"
	"net/http"
)

type Router struct {
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathHandler(w, r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to SHarepics website!!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>You may contact me at <a href=\"mailto:test@gmail.com\">myemail</a></p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	w.Write([]byte("<h1>FAQ Page</h1><p>All help you need can be found here</p>"))
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	/*if r.URL.Path == "/" {
		homeHandler(w, r)
	} else if r.URL.Path == "/contact" {
		contactHandler(w, r)
	} else {
		//w.WriteHeader(404)
		//w.Write([]byte("Page not found"))
		http.Error(w, "Page not found", 404)
	}*/
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found", 404)
	}
}

func main() {
	//http.ListenAndServe(":2020", &Router{})
	//http.ListenAndServe(":2020", http.HandlerFunc(pathHandler))
	var router http.HandlerFunc
	router = pathHandler
	http.ListenAndServe(":2020", router)
}
