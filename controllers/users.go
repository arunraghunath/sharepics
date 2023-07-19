package controllers

import (
	"fmt"
	"net/http"

	"github.com/arunraghunath/sharepics/models"
)

type Templates struct {
	New Template
}
type User struct {
	Templates
	UserService *models.UserService
}

func (u User) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u User) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error while parsing the request", http.StatusBadRequest)
		return
	}
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "user created %v", user)
}
