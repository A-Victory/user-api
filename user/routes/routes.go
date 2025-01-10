package routes

import (
	"net/http"

	"github.com/A-Victory/user-mig/user/routes/handler"
	"github.com/A-Victory/user-mig/user/service"
	"github.com/go-chi/chi/v5"
)

func Router(s *service.Service) *chi.Mux {

	router := chi.NewRouter()
	router.Use(setJSONContentType)

	h := handler.NewHandlers(s)

	router.Route("/api", func(r chi.Router) {
		r.Post("/signup", h.SignUp)
		r.Post("/signin", h.SignIn)
		r.Get("/get-users", h.ListUsers)
	})

	return router

}

func setJSONContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
