package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger	
}


func NewHello(l *log.Logger) *Hello{
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// hadnler内部にログを記載するのはユニットテストのtestablilityが下がるので推奨しない
	// 後ほどログは依存注入する
	// log.Println("Hello world")	
	h.l.Println("Hello world")	
	d, err := ioutil.ReadAll(r.Body)
	if err != nil{
		http.Error(rw, "Ooops", http.StatusBadRequest)
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Ooops"))
		return
	}
	
	fmt.Fprintf(rw, "Hello %s", d)
}