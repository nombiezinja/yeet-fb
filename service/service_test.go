package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/nombiezinja/yeet-fb/internal/config"
	"github.com/stretchr/testify/assert"
)

// TestNewServiceSuccess ensures New returns a valid Service and Server
func TestNewServiceWithMockConfig(t *testing.T) {
	mockCfg := config.MockConfig{}
	r := chi.NewMux()
	svc, srv := New(&mockCfg, r)
	assert.NotNil(t, svc)
	assert.NotNil(t, srv)
	assert.Equal(t, &mockCfg, srv.Config)
}

// TestHandlerMethodsNotImplemented checks all handler stubs return 501
func TestHandlerMethodsNotImplemented(t *testing.T) {
	cfg := &config.Config{}
	r := chi.NewMux()
	_, srv := New(cfg, r)

	// Helper to test handler
	testNotImplemented := func(handler func(http.ResponseWriter, *http.Request), name string) {
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		handler(w, req)
		assert.Equal(t, http.StatusNotImplemented, w.Code, name)
	}

	testNotImplemented(srv.HealthCheck, "HealthCheck")
	testNotImplemented(srv.GetFavicon, "GetFavicon")
	testNotImplemented(srv.HandleFacebookCallback, "HandleFacebookCallback")
	testNotImplemented(srv.RefreshFacebookToken, "RefreshFacebookToken")
	testNotImplemented(srv.VerifyFacebookToken, "VerifyFacebookToken")
	testNotImplemented(srv.Login, "Login")
	testNotImplemented(srv.Logout, "Logout")
	testNotImplemented(srv.Register, "Register")
	testNotImplemented(srv.CreatePrincipal, "CreatePrincipal")
}

// TestNewServiceFailure checks New with nil config returns valid objects
func TestNewServiceFailure(t *testing.T) {
	r := chi.NewMux()
	svc, srv := New(nil, r)
	assert.NotNil(t, svc)
	assert.NotNil(t, srv)
	assert.Nil(t, srv.Config)
}
