package main

import (
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/health"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/logger"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/server"
	"github.com/hscasn/modelmonorepo/modules/main/config"
	"github.com/hscasn/modelmonorepo/modules/main/internal/api"
)

func main() {
	config := config.New()
	log := logger.New(config.Name, false)

	onClose := func() {
		log.Infof("Server %s is shutting down\n", config.Name)
	}

	healthChecks := health.Checks{}

	srv := server.New(log, healthChecks, config.Port, onClose)
	api.New(log, srv.Router, config.UsersServiceURL)
	srv.Start()
}
