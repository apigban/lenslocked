package controllers

import (
	"net/http"

	"github.com/apigban/lenslocked/views"
)

type Users struct {
	Templates struct {
		New views.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	//View to render
	u.Templates.New.Execute(w, nil)
}
