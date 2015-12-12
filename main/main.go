package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/uprn/{uprn}", UprnHandler)
	http.Handle("/", r)
}

//HomeHandler responds to requests from the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not yet implemented", http.StatusNotImplemented)
}

//UprnHandler responds to requests for property information
func UprnHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not yet implemented", http.StatusNotImplemented)
}
