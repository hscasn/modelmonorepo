package api

import (
	"fmt"
	"net/http"

	"github.com/hscasn/modelmonorepo/lib/go/httpserver/apihandlers"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/logger"

	"github.com/go-chi/chi"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/apiresponse"
	"github.com/hscasn/modelmonorepo/lib/go/network"
	"github.com/hscasn/modelmonorepo/modules/main/pkg/api"
	usersAPI "github.com/hscasn/modelmonorepo/modules/users/pkg/api"
)

// New will bind this API to an exiting router
func New(log logger.Interface, router chi.Router, usersServiceURL string) {
	router.Post("/v1", defaultHandler(log))
	router.Post("/v1/allUsers", getAllUsersHandler(log, usersServiceURL))
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

func getAllUsersHandler(log logger.Interface, usersServiceURL string) http.HandlerFunc {
	body := api.RequestAllUsers{}
	return apihandlers.SimpleHandler(log, &body, func() (*apiresponse.ResponseData, error) {
		req := usersAPI.RequestAllUsers{}
		res := usersAPI.ResponseAllUsers{}
		err := network.PostCloudRunCall(usersServiceURL, "v1/allUsers", req, &res)
		if err != nil {
			return nil, fmt.Errorf("external api call failed: %w", err)
		}
		result := api.ResponseAllUsers{
			Users: make(api.User, len(res.Users)),
		}
		for i, user := range res.Users {
			result.Users[i] = api.User{name: user.Name}
		}
		return &apiresponse.ResponseData{
			Result: result,
		}, nil
	})
}
