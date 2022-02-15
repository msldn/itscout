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

// CI Model
type Ci struct {
	ID   string `json: "id"`
	Type string `json: "type"`
	Name string `json: "name"`
}

// Inint CIs Array
var CIs []Ci

// Index handler
func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ITScout is a discovery based CMDB")
}

// Health Check
func healthz_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This service is health")
}

// Get CIs
func getCis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CIs)
}

// Get one CI
func getCi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range CIs {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Ci{})
}

// Create New CI
func createCi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ci Ci
	_ = json.NewDecoder(r.Body).Decode(&ci)
	ci.ID = strconv.Itoa(rand.Intn(99)) // generate random id
	CIs = append(CIs, ci)
	json.NewEncoder(w).Encode(ci)

}

// Update existing CI
func updateCi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range CIs {
		if item.ID == params["id"] {
			CIs = append(CIs[:index], CIs[index+1:]...)
			var ci Ci
			_ = json.NewDecoder(r.Body).Decode(&ci)
			ci.ID = params["id"] // generate random id
			CIs = append(CIs, ci)
			json.NewEncoder(w).Encode(CIs)
			break
		}
	}

}

// Delete existing CI
func deleteCi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range CIs {
		if item.ID == params["id"] {
			CIs = append(CIs[:index], CIs[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(CIs)
}

func main() {

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/healthz/", healthz_handler)

	CIs = append(CIs, Ci{ID: "11", Type: "Server", Name: "FirstServer"})
	CIs = append(CIs, Ci{ID: "12", Type: "Server", Name: "SecondServer"})
	CIs = append(CIs, Ci{ID: "13", Type: "Server", Name: "ThirdServer"})
	CIs = append(CIs, Ci{ID: "21", Type: "Application", Name: "FirstApp"})
	CIs = append(CIs, Ci{ID: "22", Type: "Application", Name: "SecondApp"})
	CIs = append(CIs, Ci{ID: "23", Type: "Application", Name: "ThirdApp"})

	r := mux.NewRouter()

	r.HandleFunc("/api/cis", getCis).Methods("GET")
	r.HandleFunc("/api/cis/{id}", getCi).Methods("GET")
	r.HandleFunc("/api/cis", createCi).Methods("POST")
	r.HandleFunc("/api/cis/{id}", updateCi).Methods("PUT")
	r.HandleFunc("/api/cis/{id}", deleteCi).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
