package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sample-go-nakagome/handlers"
	"time"

	"github.com/gorilla/mux"
)

// v5  以下のmodifications
// 1.gorilla/mux router framework書き換え
func main() {

	// http server
	l := log.New(os.Stdout, "sample-api-1", log.LstdFlags)

	// handler作成
	ph := handlers.NewProducts(l)
	// serve mux(route router)作成し、handlerを登録
	sm := mux.NewRouter()
	// serve mux(sub router)を作成し、handlerを登録
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	// getRouter.HandleFunc("/products", ph.GetProducts)
	getRouter.HandleFunc("/products", handlers.CorsHandler(ph.GetProducts))
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	
	// sm.Handle("/products", ph)
	
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
			l.Fatal(err)
		}
	}()
	
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <- sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	// reliability, canary deploymentなどの時に重要になってくる
	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	s.Shutdown(tc)
}