package api

import (
	"github.com/go-chi/chi"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/api/health"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/api/ready"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/api/settings"
	healthPkg "github.com/hscasn/modelmonorepo/lib/go/httpserver/health"
)

// New will bind this API to an existing router
func New(router *chi.Mux, healthChecks healthPkg.Checks) {
	router.Route("/_internal/health", func(r chi.Router) {
		health.New(r, healthChecks)
	})
	router.Route("/_internal/ready", func(r chi.Router) {
		ready.New(r, healthChecks)
	})
	router.Route("/_internal/settings", func(r chi.Router) {
		settings.New(r)
	})
}
