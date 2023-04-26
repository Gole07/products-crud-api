package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Product struct (Model)
type Product struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Short_desc   string `json:"short_desc"`
	Descrtiption string `json:"description"`
	Price        int64  `json:"price"`
	Create       string `json:"create"`
}

// Init products var as slice Product struct
var products []Product

// Get All Products
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// Get Single Product
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// loop thought p roducts and find with id
	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}

// Create a new Book
func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = strconv.Itoa(rand.Intn(10000000))
	products = append(products, product)
	json.NewEncoder(w).Encode(product)

}

// Update Book
func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			var product Product
			_ = json.NewDecoder(r.Body).Decode(&product)
			product.ID = strconv.Itoa(rand.Intn(10000000))
			products = append(products, product)
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	json.NewEncoder(w).Encode(products)
}

// Delete Book
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(products)
}

func main() {
	//Init Router
	r := mux.NewRouter()
	fmt.Println("Starting server at port 7500")
	products = append(products, Product{ID: "1", Name: "Tesla Model 3", Short_desc: "T M3", Descrtiption: "Tesla Model 3 is electric car", Price: 70000, Create: "10.05.2020"})
	products = append(products, Product{ID: "2", Name: "Audi", Short_desc: "A Q8", Descrtiption: "Audi Model Q8 is not electric car", Price: 80000, Create: "19.05.2022"})
	products = append(products, Product{ID: "3", Name: "Lamborghini", Short_desc: "L Aventador", Descrtiption: "Lamborghini Aventador is electric car", Price: 300000, Create: "19.07.2023"})
	//Router Handlers/ Endpoints
	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/api/products", createProduct).Methods("POST")
	r.HandleFunc("/api/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/api/products/{id}", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":7500", r))
}
