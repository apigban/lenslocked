package main

import (
	"fmt"
	"net/http"

	"github.com/apigban/lenslocked/controllers"
	"github.com/apigban/lenslocked/models"
	"github.com/apigban/lenslocked/templates"
	"github.com/apigban/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := models.UserService{
		DB: db,
	}

	usersC := controllers.Users{
		UserService: &userService, // UserService above is not a pointer - zero value needs to be nil when not set
	}

	usersC.Templates.New = views.Must(
		views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))

	usersC.Templates.SignIn = views.Must(
		views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))

	r.Get("/signin", usersC.SignIn)
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
