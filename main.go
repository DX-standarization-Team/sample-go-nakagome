package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// GO httpパッケージのMUXにpathを登録するメソッド
	// MUX：Httpハンドラーデフォルトサーバ
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		log.Println("hello")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil{
			http.Error(rw, "Ooops", http.StatusBadRequest)
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Ooops"))
		return
		}
	
		fmt.Fprintf(rw, "Hello %s", d)
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