package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

const tmpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Number}}, {{.Road}} - OStute</title>
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
		Road:     dpa.THOROUGHFARENAME,
		Town:     dpa.POSTTOWN,
		Postcode: dpa.POSTCODE,
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
