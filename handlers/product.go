package handlers

import (
	"log"
	"net/http"

	"github.com/hochitai/microservice/data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		p.getProducts(rw, r)
	}


	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}