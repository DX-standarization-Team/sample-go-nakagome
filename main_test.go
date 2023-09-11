package main

import (
	"net/http"
	"sample-go-nakagome/handlers"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func testClient(t *testing.T){
	handler := handlers.NewApiServer()
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", handler.GetProducts)
	http.ListenAndServe(":9090", sm)
	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			// l.Fatal(err)
			println(err)
		}
	}()
}