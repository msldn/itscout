package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// CI Model
type Ci struct {
	ID         int    `json: "id"`
	Type       string `json: "type"`
	Name       string `json: "name"`
	Created_on string `json: "created_on"`
}

// Set up Datbase and open connection

// Index handler
func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ITScout is a discovery based CMDB")
}

// Health Check
func healthz_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This service is health")
}

// Error handler
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}

// Get CIs Handler
func getCisHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getCis())
}

// Get one CI
func getCiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var CIs []Ci
	var ids = []int{}
	params := mux.Vars(r)
	for _, i := range strings.Split(params["id"], ",") {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		ids = append(ids, j)
	}
	CIs = getCi(ids)
	if len(CIs) < 1 {
		errorHandler(w, r, http.StatusNotFound)
	} else if len(CIs) == 1 {
		json.NewEncoder(w).Encode(CIs[0])
	} else {
		json.NewEncoder(w).Encode(CIs)
	}
}

// Create New CI
func createCiHandler(w http.ResponseWriter, r *http.Request) {
	var ci Ci
	_ = json.NewDecoder(r.Body).Decode(&ci)
	var maxid int
	mxrow := exec_query("Select max(id) from cis")
	mxrow.Next()
	mxrow.Scan(&maxid)
	ci.ID = maxid + 1
	q := "Insert into cis values (" + strconv.Itoa(ci.ID) + ",'" + ci.Type + "','" + ci.Name + "', now() )"
	exec_query(q)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getCis())
}

// Update existing CI
func updateCiHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ci Ci
	_ = json.NewDecoder(r.Body).Decode(&ci)
	q := "update cis set type = " + "'" + ci.Type + "' , name = '" + ci.Name + "' where id = " + params["id"]
	exec_query(q)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getCis())
}

// Delete existing CI
func deleteCiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	q := "Delete From cis where ID in (" + params["id"] + ")"
	exec_query(q)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getCis())
}

func main() {

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/healthz/", healthz_handler)

	r := mux.NewRouter()

	r.HandleFunc("/api/cis", getCisHandler).Methods("GET")
	r.HandleFunc("/api/cis/{id}", getCiHandler).Methods("GET")
	r.HandleFunc("/api/cis", createCiHandler).Methods("POST")
	r.HandleFunc("/api/cis/{id}", updateCiHandler).Methods("PUT")
	r.HandleFunc("/api/cis/{id}", deleteCiHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
