package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

var (
	httpAddr = flag.String("http-addr", ":8080", "HTTP address")
)

func main() {
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/cheers", func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("X-Auth-Username")
		if username == "" {
			w.WriteHeader(http.StatusBadRequest)
			writeString(w, "you are not authorized")
			return
		}

		writeString(w, fmt.Sprintf("cheers %s!", username))
	})
	fmt.Println(http.ListenAndServe(*httpAddr, mux))
}

func writeString(w io.Writer, msg string) {
	if _, err := w.Write([]byte(msg)); err != nil {
		fmt.Printf("Error writing response: %v\n", err)
	}
}
