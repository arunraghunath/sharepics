package main

import (
	"fmt"
	"net/http"

	"github.com/arunraghunath/sharepics/controllers"
	"github.com/arunraghunath/sharepics/templates"
	"github.com/arunraghunath/sharepics/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

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

	tpl := views.Must(views.ParseFS(templates.FS, "layout-page.html", "home-page.html"))
	router.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "layout-page.html", "contact-page.html"))
	router.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.html"))
	router.Get("/faq", controllers.FAQ(tpl))

	router.Get("/testurlparam", urlParamHandler)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", 404)
	})

	fmt.Println("Starting the server at :2020")
	http.ListenAndServe(":2020", router)
}
