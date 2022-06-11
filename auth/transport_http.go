package auth

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func MakeHTTPHandler(endpoints Endpoints, logger log.Logger) http.Handler {
	mux := chi.NewMux()

	options := []kithttp.ServerOption{
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
		kithttp.ServerBefore(setLogger(logger)),
		kithttp.ServerErrorLogger(level.Error(logger)),
	}

	authHandler := kithttp.NewServer(
		endpoints.Auth,
		kithttp.NopRequestDecoder,
		encodeAuthResponse,
		options...,
	)

	mux.Get("/auth", authHandler.ServeHTTP)

	return mux
}

func setLogger(logger log.Logger) kithttp.RequestFunc {
	return func(ctx context.Context, r *http.Request) context.Context {
		return loggerToContext(ctx, logger)
	}
}

func encodeAuthResponse(ctx context.Context, w http.ResponseWriter, response any) error {
	resp := response.(AuthResponse)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("X-Auth-Username", "mcosta74")
	w.Header().Add("X-Auth-Token", resp.Token)
	return nil
}
