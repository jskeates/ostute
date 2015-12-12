package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//NotYetImplementedMessage is the text returned to the user when a request is
//made for a feature that isn't yet functional
const NotYetImplementedMessage = "Not yet implemented, sorry!"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/uprn/{uprn}", UprnHandler)
	http.ListenAndServe(":"+port, r)
}

//HomeHandler responds to requests from the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, NotYetImplementedMessage, http.StatusNotImplemented)
}

//UprnHandler responds to requests for property information
func UprnHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, NotYetImplementedMessage, http.StatusNotImplemented)
}
