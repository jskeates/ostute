package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
)

const tmpl = `
<!DOCTYPE html>
<html>
  <head>
     <meta charset="UTF-8">
     <title>{{.Number}}, {{.Road}} - OStute</title>
     <meta name="viewport" content="width=device-width, initial-scale=1.0">
     <style>
       body {
         background-color: #453c90;
         color: white;
         font-family: "Source Sans Pro Regular","Helvetica Neue",Helvetica,Arial,sans-serif;
       }
     </style>
   </head>
   <body>
     <h1>{{.Number}}, {{.Road}}</h1>
     <h2>{{.Town}}, {{.Postcode}}</h2>
   </body>
</html>
`

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
	data := struct {
		Number, Road, Town, Postcode string
	}{
		Number:   dpa.BUILDINGNUMBER,
		Road:     capitalise(dpa.THOROUGHFARENAME),
		Town:     capitalise(dpa.POSTTOWN),
		Postcode: strings.ToUpper(dpa.POSTCODE),
	}
	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Capitalise returns the given text with the first letter of every
//word capitalised.
func capitalise(text string) string {
	result := strings.ToLower(text)
	result = strings.Title(result)
	return result
}
