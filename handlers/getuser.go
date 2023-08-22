package handlers

import (
	"log"
	"net/http"
)

type Getuser struct {
	l *log.Logger
}

func NewGetuser(l*log.Logger) *Getuser {
	return &Getuser{l}
}

func (g*Getuser) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	rw.Write([]byte("getUser"))
}