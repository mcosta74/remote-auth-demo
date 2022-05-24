package main

import (
	"flag"
	"net/http"
	"strings"
)

var (
	httpAddr = flag.String("http-addr", ":8080", "HTTP address")
)

func main() {
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("missing authorization header"))
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid authorization header"))
			return
		}

		token := authHeader[7:]
		if token != "1122334455" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid token"))
			return
		}

		w.Header().Add("X-Auth-Username", "mcosta74")
		w.Write([]byte("authorized"))
	})

	http.ListenAndServe(*httpAddr, mux)
}
