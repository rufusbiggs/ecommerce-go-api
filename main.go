package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

var products = []Product{}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("products", getProducts).Methods("GET")
	r.HandleFunc("products/{id}", getProduct).Methods("GET")
	r.HandleFunc("products", createProduct).Methods("POST")
	r.HandleFunc("products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("products/{id}", deleteProduct).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, product := range products {
		if product.ID == params["id"] {
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	http.NotFound(w, r)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	json.NewDecoder(r.Body).Decode(&product)
	product.ID = strconv.Itoa(len(products) + 1)
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, product := range products {
		if product.ID == params["id"] {
			var updatedProduct Product
			json.NewDecoder(r.Body).Decode(&updatedProduct)
			updatedProduct.ID = product.ID
			products[i] = updatedProduct
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, product := range products {
		if product.ID == params["id"] {
			products = append(products[:i], products[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}