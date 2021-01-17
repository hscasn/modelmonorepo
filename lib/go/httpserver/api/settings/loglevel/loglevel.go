package loglevel

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/apiresponse"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/logger"
)

// New will bind this API to an exiting router
func New(router chi.Router) {
	router.Get("/", getLevel)

	router.Put("/", putLevelFor(logger.WarnLevel, true))
	router.Put("/debug", putLevelFor(logger.DebugLevel, false))
	router.Put("/info", putLevelFor(logger.InfoLevel, false))
	router.Put("/warn", putLevelFor(logger.WarnLevel, false))
	router.Put("/error", putLevelFor(logger.ErrorLevel, false))
	router.Put("/fatal", putLevelFor(logger.FatalLevel, false))
}

type readyResponse struct {
	Status string `json:"status"`
}

var usage = "Specify a level by hitting the endpoint with /debug, /info, " +
	"/warn, /error, or /fatal with PUT method"

func getLevel(w http.ResponseWriter, r *http.Request) {
	level := logger.GetLevel()
	message := fmt.Sprintf("Current level: %s", level.String())

	apiresponse.SendJSONResponse(apiresponse.ResponseData{
		Warnings: []string{usage},
		Result:   message,
	}, w)
}

func putLevelFor(level logger.Level, isDefault bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.SetLevel(level)

		message := fmt.Sprintf("Level changed to: %s", level.String())
		lvlNotSpecifiedMsg := fmt.Sprintf(
			"Level not specified; falling back to default. %s",
			usage)

		rData := apiresponse.ResponseData{
			Warnings: []string{},
			Result:   message,
		}

		if isDefault {
			rData.Warnings = append(
				rData.Warnings,
				lvlNotSpecifiedMsg)
		}

		apiresponse.SendJSONResponse(rData, w)
	}
}
