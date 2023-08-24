package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"sample-go-nakagome/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	
	rw.Header().Set("Access-Control-Allow-Headers", "*")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS")
	if r.Method == "OPTIONS" {
		rw.WriteHeader(http.StatusOK)
		return
	}

	// handle the request for a list of products
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	// mod. 1 post method
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	// mod. 3 put method
	if r.Method == http.MethodPut {
		// expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil{
			p.l.Println("Invalid URI unable to convert to number", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.l.Println("got id", id)

		p.updateProducts(id, rw, r)
		return
	}
	// catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Product")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json.", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}
func (p *Products) putProduct(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusBadRequest)
	}
}

func (p Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}