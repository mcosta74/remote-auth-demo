package auth

import (
	"context"
	"errors"
	"strings"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

type contextKey int

const (
	LoggerKey contextKey = iota
)

func bearerTokenFromContext(ctx context.Context) (string, error) {
	header := ctx.Value(kithttp.ContextKeyRequestAuthorization).(string)
	if header == "" {
		return "", errors.New("missing authorization header")
	}

	if !strings.HasPrefix(header, "Bearer ") {
		return "", errors.New("malformed authorization header")
	}
	return header[7:], nil
}

func loggerToContext(ctx context.Context, logger log.Logger) context.Context {
	return context.WithValue(ctx, LoggerKey, logger)
}

func loggerFromContext(ctx context.Context) log.Logger {
	if logger, ok := ctx.Value(LoggerKey).(log.Logger); ok {
		return logger
	}
	return log.NewNopLogger()
}
