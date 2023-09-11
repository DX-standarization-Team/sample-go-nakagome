package handlers

import (
	"net/http"
)

type CorsHandle func(rw http.ResponseWriter, r *http.Request)

func CorsHandler(handle CorsHandle) http.HandlerFunc {
    return func(rw http.ResponseWriter, r *http.Request) {
        //ヘッダの追加
        rw.Header().Set("Access-Control-Allow-Headers", "*")
        rw.Header().Set("Access-Control-Allow-Origin", "*")
        rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

        //プリフライトリクエストへの応答
        if r.Method == "OPTIONS" {
            rw.WriteHeader(http.StatusOK)
            return
        }

        //Handler関数の実行
        handle(rw, r)
    }
}