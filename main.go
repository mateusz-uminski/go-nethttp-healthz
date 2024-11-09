package main

import (
	"context"
	"os"

	"github.com/mateusz-uminski/go-nethttp-healthz/api"
	"github.com/mateusz-uminski/go-nethttp-healthz/util/config"
	"github.com/mateusz-uminski/go-nethttp-healthz/util/http"
	"github.com/mateusz-uminski/go-nethttp-healthz/util/log"
)

func main() {
	_ = context.Background()

	logger := log.Make(os.Stdout)

	conf := config.New("APP")

	healthzHandler := api.Healthz(logger)

	router := http.NewRouter()
	router.RegisterEndpoints("/api/v1", []http.Endpoint{
		{Path: "/healthz", Handler: healthzHandler},
	})

	server := http.NewServer(
		http.ServerWithAddr(conf.GetHost()),
		http.ServerWithPort(conf.GetPort()),
		http.ServerWithRouter(router),
	)

	logger.Infof("Starting http server on %s:%d", conf.GetHost(), conf.GetPort())

	server.Start()
}
