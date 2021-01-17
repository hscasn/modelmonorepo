package api

import (
	"fmt"
	"net/http"

	"github.com/hscasn/modelmonorepo/lib/go/httpserver/apihandlers"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/logger"

	"github.com/go-chi/chi"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/apiresponse"
	"github.com/hscasn/modelmonorepo/modules/_model/pkg/api"
)

// New will bind this API to an exiting router
func New(log logger.Interface, router chi.Router) {
	router.Post("/v1", defaultHandler(log))
}

func defaultHandler(log logger.Interface) http.HandlerFunc {
	body := api.RequestDefault{}
	return apihandlers.SimpleHandler(log, &body, func() (*apiresponse.ResponseData, error) {
		num := body.Number

		return &apiresponse.ResponseData{
			Result: api.ResponseDefault{
				Message: fmt.Sprintf("You sent the number: %d", num),
			},
		}, nil
	})
}
