package handlers
import (
	"log"
	"net/http"
	"github.com/johnnguyen-nodejs/golang/data"
	"github.com/gorilla/mux"
	// "regexp"
	"strconv"
	"context"
)
type Products struct {
	l *log.Logger
}
func NewProducts(l*log.Logger) * Products {
	return &Products{l}
}


// return the products from data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	//fetch the products from the datastore
	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
// add a product to data store
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")
	
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

// update product

func (p Products) UpdateProducts(rw http.ResponseWriter, r*http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to coonvert ID", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT Products", id)
	
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
}

// middleware
type KeyProduct struct{}

func (p Products) MiddlewareProductvalidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}