package service

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nombiezinja/yeet-fb/internal/config"
	"github.com/nombiezinja/yeet-fb/internal/generatedserver"
)

type Service struct {
	Handler *http.Handler
}

// New returns a new service, inject config and mux as dependencies
// to enable testing.
func New(config *config.Config, r *chi.Mux) (*Service, Server) {
	s := Server{
		Config: config,
	}

	opts := generatedserver.ChiServerOptions{
		BaseRouter: r,
	}

	//Mount generated server in uwu.gen.go
	handler := generatedserver.HandlerWithOptions(s, opts)

	return &Service{
		Handler: &handler,
	}, s
}

type Server struct {
	Config *config.Config
}

// --- Auth & OAuth ---
func (s Server) GetFacebookAuthUrl(w http.ResponseWriter, r *http.Request, params generatedserver.GetFacebookAuthUrlParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) HandleFacebookCallback(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) RefreshFacebookToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) VerifyFacebookToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) Logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) Register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// --- System ---
func (s Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) GetFavicon(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// --- Principals ---
func (s Server) ListPrincipals(w http.ResponseWriter, r *http.Request, params generatedserver.ListPrincipalsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) CreatePrincipal(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) DeletePrincipal(w http.ResponseWriter, r *http.Request, principalId generatedserver.PrincipalId) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) GetPrincipal(w http.ResponseWriter, r *http.Request, principalId generatedserver.PrincipalId) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) UpdatePrincipal(w http.ResponseWriter, r *http.Request, principalId generatedserver.PrincipalId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// --- Posts ---
func (s Server) DeleteAllPosts(w http.ResponseWriter, r *http.Request, principalId generatedserver.PrincipalId, params generatedserver.DeleteAllPostsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) ListPosts(w http.ResponseWriter, r *http.Request, principalId generatedserver.PrincipalId, params generatedserver.ListPostsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) DeletePost(w http.ResponseWriter, r *http.Request, principalId generatedserver.PrincipalId, postId generatedserver.PostId) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) GetPost(w http.ResponseWriter, r *http.Request, principalId generatedserver.PrincipalId, postId generatedserver.PostId) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) BatchDeletePosts(w http.ResponseWriter, r *http.Request, principalId generatedserver.PrincipalId) {
	w.WriteHeader(http.StatusNotImplemented)
}
