package handlers

import (
	"net/http"

	"sample-go-nakagome/data"
)

// ServerInterface を実装するような struct を定義
type ApiServer struct {
}

func NewApiServer() *ApiServer {
	return &ApiServer{}
}

func (si ApiServer) PostProduct(rw http.ResponseWriter, r *http.Request) {
}
func (si ApiServer) GetProduct(rw http.ResponseWriter, r *http.Request) {
}
func (si ApiServer) PatchProductsProductId(rw http.ResponseWriter, r *http.Request) {
}
func (si ApiServer) GetProducts(rw http.ResponseWriter, r *http.Request) {
	// レスポンスデータの作成
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
    /* 諸々の処理 */
    rw.WriteHeader(http.StatusOK)
}