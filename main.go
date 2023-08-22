package main

import (
	"log"
	"net/http"
	"os"

	"sample-go-nakagome/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	// GO httpパッケージのMUXにpathを登録するメソッド
	// MUX：Httpハンドラーデフォルトサーバ
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		log.Println("hello")	
	})
	http.HandleFunc("/getUser", func(http.ResponseWriter, *http.Request){
		log.Println("getUser")	
	})
	http.HandleFunc("/createUser", func(http.ResponseWriter, *http.Request){
		log.Println("createUser")	
	})
	// Httpサーバを構築してデフォルトハンドラを登録するメソッド
	// 第2引数がハンドラーで、指定なしだとデフォルトハンドラMUXが呼び出される
	// path - func を Mappingする
	http.ListenAndServe(":9090", nil)
}