package main

import(
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Karyawan struct {
	ID		string `json:"id"`
	Name	string `json:"name"`
	NIP		int `json:"nip"`
	Divisi	string `json:"divisi"`
	Alamat	string `json:"alamat"`
	Umur	string `json:"umur"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Selamat datang din home page")
}

func allKaryawans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(Karyawans)
}

func singleKaryawan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for _, karyawan := range Karyawans {
		if karyawan.ID ==id {
			json.NewEncoder(w).Encode(karyawan)
		}
	}
}

func createKaryawan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var karyawan Karyawan
	err := json.NewDecoder(r.Body).Decode(&karyawan)
	if err !=nil {
		fmt.Println(err.Error())
		return
	}
	Karyawans = append(Karyawans, karyawan)
	json.NewEncoder(w).Encode(karyawan)
}

func updateKaryawan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","application/json")
	params := mux.Vars(r)
	id := params["id"]

	var karyawan Karyawan
	err := json.NewDecoder(r.Body).Decode(&karyawan)
	if err !=nil {
		fmt.Println(err.Error())
		return
	}

	for i, p := range Karyawans {
		if p.ID == id {
			Karyawans[i].Name = karyawan.Name
			Karyawans[i].NIP = karyawan.NIP
			Karyawans[i].Divisi = karyawan.Divisi
			Karyawans[i].Alamat = karyawan.Alamat
			Karyawans[i].Umur = karyawan.Umur
		}
	}
}

func deleteKaryawan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id:= params["id"]

	for i, p := range Karyawans {
		if p.ID == id {
			Karyawans = append(Karyawans[:i], Karyawans[i+1:]...)
			json.NewEncoder(w).Encode(p)
		}
	}
}

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/",home)
	r.HandleFunc("/karyawans",allKaryawans).Methods("GET")
	r.HandleFunc("/karyawans/{id}", singleKaryawan).Methods("GET")
	r.HandleFunc("/karyawans", createKaryawan).Methods("POST")
	r.HandleFunc("/karyawans/{id}", updateKaryawan).Methods("PUT")
	r.HandleFunc("/karyawans/{id}", deleteKaryawan).Methods("DELETE")
	fmt.Println("Application runing")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	Karyawans = []Karyawan{
		Karyawan{
			ID: "1",
			Name: "Andi",
			NIP: 202020,
			Divisi: "IT Support",
			Alamat: "Jalan Minggu",
			Umur: "24",
		},{
			ID: "2",
			Name: "Dido",
			NIP: 202020,
			Divisi: "IT Specialist",
			Alamat: "Jalan Senin",
			Umur: "21",
		},{
			ID: "3",
			Name: "Iki",
			NIP: 202020,
			Divisi: "IT Support",
			Alamat: "Jalan Kamis",
			Umur: "22",
		},
	}

	handleRequest()
}

var Karyawans []Karyawan