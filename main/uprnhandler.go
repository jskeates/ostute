package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

//UprnHandler responds to requests for property information
func UprnHandler(w http.ResponseWriter, r *http.Request) {
	uprn := mux.Vars(r)["uprn"]
	dpa, err := FetchDpaInfo(uprn, APIKey)
	if err != nil {
		if err.Error() == ErrorUprnNotFound {
			http.Error(w, fmt.Sprintf("UPRN %s not found", uprn), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	_, err = io.WriteString(w, fmt.Sprintf("<html><body><h1>%s :: %s</h1><h2>%s :: %s</h2></body></html>", dpa.BUILDINGNUMBER, dpa.THOROUGHFARENAME, dpa.POSTTOWN, dpa.POSTCODE))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
