package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//NotYetImplementedMessage is the text returned to the user when a request is
//made for a feature that isn't yet functional
const NotYetImplementedMessage = "Not yet implemented, sorry!"

//APIKey is the API key for accessing the OS API endpoints
var APIKey string

func main() {
	port := getEnvVarOrDie("PORT")
	APIKey = getEnvVarOrDie("OS_API_KEY")

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/uprn/{uprn}", UprnHandler)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":"+port, loggedRouter)
}

//HomeHandler responds to requests from the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, NotYetImplementedMessage, http.StatusNotImplemented)
}

func getEnvVarOrDie(envVar string) string {
	value := os.Getenv(envVar)

	if value == "" {
		log.Fatalf("Environment variable $%s must be set", envVar)
	}
	return value
}
