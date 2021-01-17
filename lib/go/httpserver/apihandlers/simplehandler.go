package apihandlers

import (
	"net/http"

	"github.com/hscasn/modelmonorepo/lib/go/httpserver/apirequest"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/apiresponse"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/customerrors"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/logger"
)

// SimpleHandler is a handler that parses the request body into requestBody (make sure to pass a & to it),
// and then calls the runner function. The runner function must either return a ResponseData
// object, or an error. If the error is an instance of errors.UserFacingError, then it is returned by the API
// with the full text and status code, otherwise it is hidden as an InternalServerError
func SimpleHandler(log logger.Interface, requestBody interface{}, runner func() (*apiresponse.ResponseData, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apirequest.Parse(r, requestBody); err != nil {
			log.Errorf("Error on SimpleHandler: %v", err)
			apiresponse.SendJSONResponse(apiresponse.ResponseData{
				Code:   http.StatusBadRequest,
				Errors: []string{"Could not parse request body"},
			}, w)
			return
		}

		result, err := runner()
		if err != nil {
			if userErr, ok := err.(customerrors.UserFacingError); ok && userErr != nil {
				errCode := userErr.StatusCode()
				if errCode < 100 {
					errCode = http.StatusBadRequest
				}
				apiresponse.SendJSONResponse(apiresponse.ResponseData{
					Code:   errCode,
					Errors: []string{userErr.UserMessage()},
				}, w)
			} else {
				log.Errorf("Internal server error on SimpleHandler: %v", err)
				apiresponse.SendJSONResponse(apiresponse.ResponseData{
					Code: http.StatusInternalServerError,
				}, w)
			}
			return
		}

		apiresponse.SendJSONResponse(*result, w)
	}
}
