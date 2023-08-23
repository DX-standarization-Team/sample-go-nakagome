package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines the structure for an API product
// `json:"id"`: id にリネーム、"-"": キーを返さない、 "omniempty": 空であればキーを返さない
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

// ToJSON コンテンツのコレクションをJSONにシリアライズする
// NewEncoder は json.Unmarshal に比べてよいパフォーマンスを提供する
// アウトプットをインメモリのバイトスライスにバッファーしなくていいから
// サービスのオーバーヘッドを削減する
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJSON(w io.Writer)error{
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// mod2. create decoder
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts()Products{
	return productList
}
func AddProduct(p *Product){
	p.ID = getNextID()
	productList = append(productList, p)
}
func getNextID() int {
	lp := productList[len(productList) -1]
	return lp.ID + 1
}
func UpdateProduct(id int, p*Product) error {
	lp := productList[len(productList) -1]
	return lp.ID + 1
}
func findProduct() (*Product, error) {
	for_, p := range productList{
		f p.ID == id {
			return p, nil
		}
	}
}

// productList is a hard coded list of products for this
// example data source
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}