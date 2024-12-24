package app

import (
	"net/http"

	"github.com/Alkestian/api-practice/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func loadRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
	})

	r.Route("/users", handleUserRoutes)

	return r
}

func handleUserRoutes(r chi.Router) {
	userHandler := &handler.User{}

	r.Get("/", userHandler.ListUsers)
	r.Post("/", userHandler.CreateUser)
	r.Get("/{id}", userHandler.GetUserByID)
	r.Put("/{id}", userHandler.UpdateUserByID)
	r.Delete("/{id}", userHandler.DeleteUserByID)
	r.Patch("/{id}/{field}", userHandler.UpdateUserField)
}