package controllers

import (
	"net/http"

	controllers "github.com/coolbambook/controllers1"
)

type Users struct {
	Templates struct {
		New controllers.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}
