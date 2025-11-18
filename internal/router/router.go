package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/nombiezinja/yeet-fb/internal/config"
)

func NewRouter(cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()

	r.NotFound(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))

	// TODO mount middlewares.

	return r
}
