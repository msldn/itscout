package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ITScout is a discovery based CMDB")
}

func healthz_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This service is health")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/healthz/", healthz_handler)
	http.ListenAndServe(":8000", nil)
}
