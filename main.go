package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"sample-go-nakagome/handlers"
)

// v3 以下のmodifications
// 1. Productsデータ用意
// 2. Products Handler作成
func main() {

	l := log.New(os.Stdout, "sample-api-1", log.LstdFlags)
	// mod. 2
	ph := handlers.NewProducts(l)
	sm := http.NewServeMux()
	sm.Handle("/", ph)
	
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