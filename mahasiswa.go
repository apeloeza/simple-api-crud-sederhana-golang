package main

import(
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
)

//struktur data
//Mahasiswa bolds your Mahasiswa attribute
type Mahasiswa struct {
	ID		 string `json:"id"`
	Title	 string `json:"title"`
	Price	 int 	`json:"price"`
	Quantity int	`json:"quantity"`
}

func home(w http.ResponseWriter, r *http.Request) { // fungsi menerima http.ResponseWriter dan Request
	fmt.Fprintf(w, "selamat datang di home page") // print out
}

//handler yg akan melempar sepua data ke ResponseWriter
func allmahasiswas(w http.ResponseWriter, r *http.Request){
	// merubah data text mennjadi json
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(mahasiswas)
}

// membuat handler untuk memfilter data sesuai dengan parameter id yang dikirimkan
func singleMahasiswa(w http.ResponseWriter, r *http.Request) {
	// merubah data text mennjadi json
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for _, Mahasiswa := range mahasiswas {
		if Mahasiswa.ID == id {
			json.NewEncoder(w).Encode(Mahasiswa)
		}
	}
}

// create
// membuat handler
func createMahasiswa(w http.ResponseWriter, r *http.Request) {
	// merubah data text mennjadi json
	w.Header().Set("content-type", "application/json")
	var Mahasiswa Mahasiswa
	// decode request budy ke struktur data produk
	err := json.NewDecoder(r.Body).Decode(&Mahasiswa)
	if err !=nil {
		fmt.Println(err.Error())
		return
	}
	// tambah koleksi data variabel Mahasiswa
	// encode
	mahasiswas = append(mahasiswas, Mahasiswa)
	json.NewEncoder(w).Encode(Mahasiswa) // (mahasiswas) --> untuk melihat semua data
}

// update
// membuat handelr mengambil parameter id dari url
func updateMahasiswa(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var Mahasiswa Mahasiswa
	// handler ini mendecode request body ke structure data Mahasiswa
	err := json.NewDecoder(r.Body).Decode(&Mahasiswa)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// mencari data dengan id tsb
	// ubah setiap valuenya
	for i, p := range mahasiswas {
		if p.ID == id {
			mahasiswas[i].Title = Mahasiswa.Title
			mahasiswas[i].Price = Mahasiswa.Price
			mahasiswas[i].Quantity = Mahasiswa.Quantity
			json.NewEncoder(w).Encode(mahasiswas[i])
			return
		}
	}
 }

 // delete
 func deleteMahasiswa(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for i, p := range mahasiswas {
		if p.ID == id {
			mahasiswas = append(mahasiswas[:i], mahasiswas[i+1:]...)
			json.NewEncoder(w).Encode(p)
			return
		}
	} 
 }


// fungsi handleRequest berfungsi sebagai route
func handleRequest(){ 
	// kita gunakan gorrila mux
	r := mux.NewRouter().StrictSlash(true)
	// HandleFunc mendaftarkan fungsi home untuk menangani url dengan pola slash "/"
	r.HandleFunc("/", home) 
	// Route View
	r.HandleFunc("/mahasiswas", allmahasiswas).Methods("GET")
	r.HandleFunc("/mahasiswas/{id}", singleMahasiswa).Methods("GET")
	// route Create
	r.HandleFunc("/mahasiswas", createMahasiswa).Methods("POST") // definisi method
	// route update
	r.HandleFunc("/mahasiswas/{id}", updateMahasiswa).Methods("PUT")
	// route Delete
	r.HandleFunc("/mahasiswas/{id}", deleteMahasiswa).Methods("DELETE")

	fmt.Println("Application running")
	// dijalankan pada port 8080
	log.Fatal(http.ListenAndServe(":8080", r)) 
}

func main(){ //fungsi main
	//dummy data
	mahasiswas = []Mahasiswa{
		Mahasiswa{ID: "1", Title: "First Mahasiswa", Price: 200000, Quantity: 5},
		Mahasiswa{ID: "2",Title: "Second Mahasiswa", Price: 500000, Quantity: 15},
	}


	handleRequest()
}

// global variabel - pengganti database
var mahasiswas []Mahasiswa