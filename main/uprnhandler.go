package main

import "net/http"

//UprnHandler responds to requests for property information
func UprnHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, NotYetImplementedMessage, http.StatusNotImplemented)
}
