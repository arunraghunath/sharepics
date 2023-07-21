package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/arunraghunath/sharepics/controllers"
	"github.com/arunraghunath/sharepics/models"
	"github.com/arunraghunath/sharepics/templates"
	"github.com/arunraghunath/sharepics/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
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

	cfg := models.DefaultPostgresConfig()
	db, err := sql.Open("postgres", cfg.String())
	if err != nil {
		panic(err)
	}
	userService := &models.UserService{
		DB: db,
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	tpl := views.Must(views.ParseFS(templates.FS, "home.html", "tailwind.html"))
	router.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.html", "tailwind.html"))
	router.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.html", "tailwind.html"))
	router.Get("/faq", controllers.FAQ(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "signup.html", "tailwind.html"))
	user := controllers.User{Templates: controllers.Templates{New: tpl}, UserService: userService}
	router.Get("/signup", user.New)
	router.Post("/signup", user.Create)

	user.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "signin.html", "tailwind.html"))
	router.Get("/signin", user.SignIn)
	router.Post("/signin", user.ProcessSignIn)

	router.Get("/users/me", user.CurrentUser)

	router.Get("/testurlparam", urlParamHandler)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", 404)
	})

	fmt.Println("Starting the server at :2020")
	http.ListenAndServe(":2020", router)
}
