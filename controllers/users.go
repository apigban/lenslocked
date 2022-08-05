package controllers

import (
	"fmt"
	"net/http"

	"github.com/apigban/lenslocked/models"
)

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, r, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		//Log the error somewhere
		//TODO: figure out better logging options
		fmt.Println(err)
		//Print a generic error for the user, so as not to leak any sensitive information
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User Created: %+v", user)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	u.Templates.SignIn.Execute(w, r, data)
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")

	//Process customer login by validating the information
	user, err := u.UserService.Authenticate(data.Email, data.Password)
	if err != nil {
		//Log the error somewhere
		//TODO: figure out better logging options
		fmt.Println(err)
		//Print a generic error for the user, so as not to leak any sensitive information
		http.Error(w, "Invalid Login.", http.StatusInternalServerError)
		return
	}

	//TODO - Implement cookie for sessions management
	cookie := http.Cookie{
		Name:     "email",
		Value:    user.Email,
		Path:     "/", // cookie can be accessed by other pages, no constraints
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "User authenticated: %+v", user)
	// TODO - redirect user to their gallery on successful login
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	email, err := r.Cookie("email")

	// TODO - redirect user to the login page if cookie is missing
	if err != nil {
		fmt.Fprint(w, "the email cookie could not be read.")
		return
	}

	fmt.Fprintf(w, "Email cookie: %s\n", email.Value)
}
