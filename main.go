package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	localAddress := fmt.Sprintf("%s:%s", os.Getenv("ADDR"), os.Getenv("PORT"))

	root := weaver.Init(context.Background())
	opts := weaver.ListenerOptions{LocalAddress: localAddress}
	lis, err := root.Listener("simple-auth-sw", opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Start listening on %s", localAddress)

	loginHandler, err := weaver.Get[LoginHandler](root)
	if err != nil {
		log.Fatal(err)
	}

	registerHandler, err := weaver.Get[RegisterHandler](root)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Method not allowed"}`))
			return
		}

		var userPayload UserPayload
		err := json.NewDecoder(r.Body).Decode(&userPayload)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
			return
		}

		token, err := loginHandler.Login(r.Context(), userPayload)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error": "Method not allowed"}`))
			return
		}

		var userPayload UserPayload
		err := json.NewDecoder(r.Body).Decode(&userPayload)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
			return
		}

		err = registerHandler.Register(r.Context(), userPayload)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	http.Serve(lis, nil)
}
