package main

import (
	"flag"
	"fmt"
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
			w.Write([]byte("you are not authorized"))
			return
		}

		w.Write([]byte(fmt.Sprintf("cheers %s!", username)))
	})
	http.ListenAndServe(*httpAddr, mux)
}
