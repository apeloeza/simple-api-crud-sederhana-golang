package main

import(
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
)

//struktur data
//product bolds your product attribute
type Product struct {
	ID		 string `json:"id"`
	Title	 string `json:"title"`
	Price	 int 	`json:"price"`
	Quantity int	`json:"quantity"`
}

func home(w http.ResponseWriter, r *http.Request) { // fungsi menerima http.ResponseWriter dan Request
	fmt.Fprintf(w, "selamat datang di home page") // print out
}

//handler yg akan melempar sepua data ke ResponseWriter
func allProducts(w http.ResponseWriter, r *http.Request){
	// merubah data text mennjadi json
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(Products)
}

// membuat handler untuk memfilter data sesuai dengan parameter id yang dikirimkan
func singleProduct(w http.ResponseWriter, r *http.Request) {
	// merubah data text mennjadi json
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for _, product := range Products {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
		}
	}
}

// create
// membuat handler
func createProduct(w http.ResponseWriter, r *http.Request) {
	// merubah data text mennjadi json
	w.Header().Set("content-type", "application/json")
	var product Product
	// decode request budy ke struktur data produk
	err := json.NewDecoder(r.Body).Decode(&product)
	if err !=nil {
		fmt.Println(err.Error())
		return
	}
	// tambah koleksi data variabel product
	// encode
	Products = append(Products, product)
	json.NewEncoder(w).Encode(product) // (Products) --> untuk melihat semua data
}

// update
// membuat handelr mengambil parameter id dari url
func updateProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var product Product
	// handler ini mendecode request body ke structure data product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// mencari data dengan id tsb
	// ubah setiap valuenya
	for i, p := range Products {
		if p.ID == id {
			Products[i].Title = product.Title
			Products[i].Price = product.Price
			Products[i].Quantity = product.Quantity
			json.NewEncoder(w).Encode(Products[i])
			return
		}
	}
 }

 // delete
 func deleteProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for i, p := range Products {
		if p.ID == id {
			Products = append(Products[:i], Products[i+1:]...)
			json.NewEncoder(w).Encode(p)
			return
		}
	} 
 }


// fungsi handleRequest sebagai route
func handleRequest(){ 
	// kita gunakan gorrila mux
	r := mux.NewRouter().StrictSlash(true)
	// HandleFunc mendaftarkan fungsi home untuk menangani url dengan pola slash "/"
	r.HandleFunc("/", home) 
	// Route View
	r.HandleFunc("/products", allProducts).Methods("GET")
	r.HandleFunc("/products/{id}", singleProduct).Methods("GET")
	// route Create
	r.HandleFunc("/products", createProduct).Methods("POST") // definisi method
	// route update
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	// route Delete
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

	fmt.Println("Application running")
	// dijalankan pada port 8080
	log.Fatal(http.ListenAndServe(":8080", r)) 
}

func main(){ //fungsi main
	//dummy data
	Products = []Product{
		Product{ID: "1", Title: "First product", Price: 200000, Quantity: 5},
		Product{ID: "2",Title: "Second Product", Price: 500000, Quantity: 15},
	}


	handleRequest()
}

// global variabel - pengganti database
var Products []Product