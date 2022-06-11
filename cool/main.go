package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"aidanwoods.dev/go-paseto"
)

var (
	httpAddr  = flag.String("http-addr", ":8080", "HTTP address")
	secretKey = os.Getenv("SECRET_KEY")
)

type User struct {
	ID        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}

func main() {
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/cheers", func(w http.ResponseWriter, r *http.Request) {
		user, err := getUserInfo(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			writeString(w, err.Error())
			return
		}

		data, err := json.MarshalIndent(user, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeString(w, err.Error())
			return
		}
		writeString(w, fmt.Sprintf("Hello\n\n%s\n", string(data)))
	})
	fmt.Println(http.ListenAndServe(*httpAddr, mux))
}

func writeString(w io.Writer, msg string) {
	if _, err := w.Write([]byte(msg)); err != nil {
		fmt.Printf("Error writing response: %v\n", err)
	}
}

func getUserInfo(r *http.Request) (*User, error) {
	authHeader := r.Header.Get("X-Auth-Token")
	if authHeader == "" {
		return nil, errors.New("not authorized")
	}

	key, err := paseto.V4SymmetricKeyFromHex(secretKey)
	if err != nil {
		return nil, err
	}

	parser := paseto.NewParser()
	parser.AddRule(paseto.IssuedBy("com.mcosta74.auth"))
	parser.AddRule(paseto.ValidAt(time.Now()))
	token, err := parser.ParseV4Local(key, authHeader, nil)
	if err != nil {
		return nil, err
	}

	userStr, err := token.GetString("userInfo")
	if err != nil {
		return nil, err
	}

	var user User
	if err := json.Unmarshal([]byte(userStr), &user); err != nil {
		return nil, err
	}
	return &user, nil
}
