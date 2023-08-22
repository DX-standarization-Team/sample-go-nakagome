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

// v2 以下のmodifications
// 1. ログの集約化
// 2. handlersのフォルダ分離
// 3．server設定（タイムアウト等）（https://pkg.go.dev/net/http#Server）
// 4. graceful shutdown
func main() {

	// mod. 1
	l := log.New(os.Stdout, "sample-api-1", log.LstdFlags)
	// mod. 2
	hh := handlers.NewHello(l)
	guh := handlers.NewGetuser(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/getUser", guh)
	
	// mod. 3
	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	
	// Httpサーバを構築してデフォルトハンドラを登録するメソッド
	// 第2引数がハンドラーに自作ハンドラーを指定する
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	
	// mod. 4
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <- sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	// reliability, canary deploymentなどの時に重要になってくる
	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	s.Shutdown(tc)
}