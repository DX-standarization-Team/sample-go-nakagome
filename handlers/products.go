package handlers

type Products {
	l *log.Logger
}

func NewProducts(l*log.Logger) * Products {
	return &Products{l}
}

func(p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	rw.Write([]byte("Products"))
}