package auth

import (
	"flag"
	"os"
)

type Config struct {
	HttpAddr  string
	LogLevel  string
	SecretKey string
}

func GetConfig(name string) (Config, error) {
	fs := flag.NewFlagSet(name, flag.ExitOnError)

	config := Config{
		LogLevel: "info",
	}

	fs.StringVar(&config.HttpAddr, "http-addr", ":8080", "HTTP address")

	config.LogLevel = os.Getenv("LOG_LEVEL")
	config.SecretKey = os.Getenv("SECRET_KEY")

	return config, nil
}
