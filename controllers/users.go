package controllers

import (
	"fmt"
	"net/http"
)

type Templates struct {
	New Template
}
type User struct {
	Templates
}

func (u User) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}

func (u User) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error while parsing the request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Email Address: %s", r.PostForm.Get("email"))
	fmt.Fprintf(w, "Password: %s", r.PostForm.Get("password"))
}
