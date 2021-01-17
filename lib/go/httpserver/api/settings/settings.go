package settings

import (
	"github.com/go-chi/chi"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/api/settings/loglevel"
)

// Create will bind this API to an exiting router
func New(router chi.Router) {
	router.Route("/loglevel", func(r chi.Router) {
		loglevel.New(r)
	})
}
