package api

import (
	"net/http"

	"github.com/hscasn/modelmonorepo/lib/go/httpserver/apihandlers"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/logger"

	"github.com/go-chi/chi"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/apiresponse"
	"github.com/hscasn/modelmonorepo/modules/users/pkg/api"
)

// New will bind this API to an exiting router
func New(log logger.Interface, router chi.Router) {
	router.Post("/v1/allUsers", defaultHandler(log))
}

func defaultHandler(log logger.Interface) http.HandlerFunc {
	body := api.RequestGetUsers{}
	return apihandlers.SimpleHandler(log, &body, func() (*apiresponse.ResponseData, error) {
		return &apiresponse.ResponseData{
			Result: api.ResponseGetUsers{
				Users: []api.User{
					api.User{Name: "Foo"},
					api.User{Name: "Bar"},
				},
			},
		}, nil
	})
}
