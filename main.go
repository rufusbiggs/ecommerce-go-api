package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
)

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

var db *sql.DB

func initDB() {
	var err error
	// Build connection string using environment variables
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	fmt.Printf("Database connection string: %s\n", connStr)

	// Retry logic to wait for database to be ready
	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			fmt.Printf("Error opening database: %v. Retrying...\n", err)
			time.Sleep(2 * time.Second)
			continue
		}
		
		pingErr := db.Ping()
		if pingErr != nil {
			log.Printf("Error pinging database: %v. Retrying...\n", pingErr)
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Println("Database connection established")
		return
	}
	
	log.Fatalf("Could not connect to the database after retries: %v", err)
}

func main() {
	initDB()
	defer db.Close()

	fmt.Println("Starting the server...")
	r := mux.NewRouter()

	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	rows, err := db.Query("SELECT id, name, description, price, stock FROM products")
	if err != nil {
		http.Error(w, "Failed to query products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock)
		if err != nil {
			http.Error(w, "Failed to scan product", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var product Product
	err := db.QueryRow("SELECT id, name, description, price, stock FROM products WHERE id = $1", id).
		Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, "Failed to query product", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}
	err = db.QueryRow(
		"INSERT INTO products (name, description, price, stock) VALUES ($1, $2, $3, $4) RETURNING id", 
		product.Name, product.Description, product.Price, product.Stock,
	).Scan(&product.ID)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err = db.Exec(
		"UPDATE products SET name=$1, description=$2, price=$3, stock=$4 WHERE id=$5",
		product.Name, product.Description, product.Price, product.Stock, id,
	)
	if err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := db.Exec("DELETE FROM products WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}