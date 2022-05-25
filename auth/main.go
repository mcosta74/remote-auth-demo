package main

import (
	"flag"
	"fmt"
	"io"
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
			writeString(w, "missing authorization header")
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer") {
			w.WriteHeader(http.StatusUnauthorized)
			writeString(w, "invalid authorization header")
			return
		}

		token := authHeader[7:]
		if token != "1122334455" {
			w.WriteHeader(http.StatusUnauthorized)
			writeString(w, "invalid token")
			return
		}

		w.Header().Add("X-Auth-Username", "mcosta74")
		writeString(w, "authorized")
	})

	fmt.Println(http.ListenAndServe(*httpAddr, mux))
}

func writeString(w io.Writer, msg string) {
	if _, err := w.Write([]byte(msg)); err != nil {
		fmt.Printf("Error writing response: %v\n", err)
	}
}
