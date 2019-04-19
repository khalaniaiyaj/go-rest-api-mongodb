package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	. "github.com/user/go-rest-api-mongodb/controller"
	. "github.com/user/go-rest-api-mongodb/middlewares"
	"log"
	"net/http"
)

var controller = Controller{}
var middleware = Middleware{}

func AllowOriginFunc(r *http.Request, origin string) bool {
		return true

}

func main() {
	controller.Init()
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
    AllowedOrigins:   []string{"*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: true,
    MaxAge:           300,
  })
	r.Use(cors.Handler)
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
