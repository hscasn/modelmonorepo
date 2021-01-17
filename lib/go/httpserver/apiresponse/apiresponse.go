package apiresponse

import (
	"encoding/json"
	"net/http"
	"time"
)

type statusAdditionalInfo struct {
	Errors   []string `json:"errors"`
	Warnings []string `json:"warnings"`
}

type status struct {
	Message        string                `json:"message"`
	AdditionalInfo *statusAdditionalInfo `json:"additionalInfo"`
}

// Response is the final response created by the model. This gets sent to
// the client as the response
type Response struct {
	ResponseNoResult
	Result interface{} `json:"result"`
}

// ResponseNoResult is the final response created by this model, but does not
// include the result. It is meant to be used to compose other types
type ResponseNoResult struct {
	Status          status `json:"status"`
	ServerTimestamp string `json:"serverTimestamp"`
	headers         map[string]string
}

// ResponseData is the input we use to build the Response. It has only
// the necessary data the factory needs
type ResponseData struct {
	Code     int
	Result   interface{}
	Headers  map[string]string
	Warnings []string
	Errors   []string
}

// SendJSONResponse accepts a ResponseData and sends an Response to the
// client
func SendJSONResponse(d ResponseData, w http.ResponseWriter) {
	// Adding default values if not present
	if d.Code == 0 {
		d.Code = 200
	}
	if d.Result == nil {
		d.Result = ""
	}

	// Setting headers
	for headerKey, headerValue := range d.Headers {
		w.Header().Set(headerKey, headerValue)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(d.Code)

	// Creating "additional info"
	var additionalInfo *statusAdditionalInfo
	if len(d.Warnings) > 0 || len(d.Errors) > 0 {
		additionalInfo = &statusAdditionalInfo{
			Errors:   []string{},
			Warnings: []string{},
		}
		for _, warning := range d.Warnings {
			additionalInfo.Warnings = append(
				additionalInfo.Warnings,
				warning)
		}
		for _, err := range d.Errors {
			additionalInfo.Errors = append(
				additionalInfo.Errors,
				err)
		}
	}

	// Creating Response that will be sent
	r := Response{
		ResponseNoResult: ResponseNoResult{
			Status: status{
				Message:        http.StatusText(d.Code),
				AdditionalInfo: additionalInfo,
			},
			ServerTimestamp: time.Now().Format(time.RFC3339),
		},
		Result: d.Result,
	}

	e, _ := json.Marshal(r)
	w.Write(e)
}
