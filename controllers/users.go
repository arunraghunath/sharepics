package controllers

import (
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
