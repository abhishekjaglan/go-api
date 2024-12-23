package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	//"github.com/avukadin/goapi/internal/middlware"
)

func Hanlder(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		// Middleware for /account route
		router.Use(middleware.Authorization) // middleware to all requests coming to /account route

		router.Get("/coins", GetCoinBalance)
	})
}
