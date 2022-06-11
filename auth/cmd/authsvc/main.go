package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/mcosta74/auth"

	"github.com/oklog/run"
)

func main() {
	config, err := auth.GetConfig("authsvc")
	if err != nil {
		os.Exit(1)
	}

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
		logger = level.NewFilter(logger, level.Allow(level.ParseDefault(config.LogLevel, level.InfoValue())))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	}
	level.Info(logger).Log("msg", "started")
	defer level.Info(logger).Log("msg", "stopped")

	tokenBuilder, err := auth.NewPasetoBuilder(config.SecretKey)
	if err != nil {
		level.Error(logger).Log("component", "TokenBuilder", "msg", "error in NewPasetoBuilder", "err", err)
		os.Exit(1)
	}

	var (
		repo        = auth.NewRepository()
		service     = auth.NewService(repo, tokenBuilder)
		endpoints   = auth.MakeEndpoints(service)
		httpHandler = auth.MakeHTTPHandler(endpoints, log.With(logger, "component", "HTTP"))
	)

	var g run.Group
	{
		// HTTP Handler
		listener, err := net.Listen("tcp", config.HttpAddr)
		if err != nil {
			level.Error(logger).Log("component", "HTTP", "msg", "error in Listen", "err", err)
			os.Exit(1)
		}

		g.Add(func() error {
			server := &http.Server{
				Handler:      httpHandler,
				ReadTimeout:  2 * time.Second,
				WriteTimeout: 2 * time.Second,
			}

			return server.Serve(listener)
		}, func(err error) {
			listener.Close()
		})
	}
	{
		// Signal Handler
		g.Add(run.SignalHandler(context.Background(), syscall.SIGTERM, syscall.SIGINT))
	}
	err = g.Run()
	level.Info(logger).Log("msg", "exit", "err", err)
}
