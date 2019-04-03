package main

import (
	"github.com/go-chi/chi"
	. "github.com/user/golang-new/controller"
	. "github.com/user/golang-new/middlewares"
	"log"
	"net/http"
)

var controller = Controller{}
var middleware = Middleware{}

func main() {
	controller.Init()
	r := chi.NewRouter()
	r.With(middleware.AuthMiddleware).Get("/movies", controller.GetAllMovies)
	r.With(middleware.AuthMiddleware).Get("/movie/{id}", controller.GetMovieById)
	r.With(middleware.AuthMiddleware).Post("/movie", controller.AddMovie)
	r.With(middleware.AuthMiddleware).Put("/movie", controller.UpdateMovie)
	r.With(middleware.AuthMiddleware).Delete("/movie/{id}", controller.DeleteMovie)
	r.HandleFunc("/login", controller.Login)
	r.HandleFunc("/register", controller.CreateUser)
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatal(err)
	}
}
