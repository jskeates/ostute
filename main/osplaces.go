package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Dpa is a struct of information about a single place
type Dpa struct {
	ADDRESS                       string  `json:"ADDRESS"`
	BLPUSTATECODE                 string  `json:"BLPU_STATE_CODE"`
	BLPUSTATECODEDESCRIPTION      string  `json:"BLPU_STATE_CODE_DESCRIPTION"`
	BLPUSTATEDATE                 string  `json:"BLPU_STATE_DATE"`
	BUILDINGNUMBER                string  `json:"BUILDING_NUMBER"`
	CLASSIFICATIONCODE            string  `json:"CLASSIFICATION_CODE"`
	CLASSIFICATIONCODEDESCRIPTION string  `json:"CLASSIFICATION_CODE_DESCRIPTION"`
	ENTRYDATE                     string  `json:"ENTRY_DATE"`
	LANGUAGE                      string  `json:"LANGUAGE"`
	LASTUPDATEDATE                string  `json:"LAST_UPDATE_DATE"`
	LOCALCUSTODIANCODE            float64 `json:"LOCAL_CUSTODIAN_CODE"`
	LOCALCUSTODIANCODEDESCRIPTION string  `json:"LOCAL_CUSTODIAN_CODE_DESCRIPTION"`
	LOGICALSTATUSCODE             string  `json:"LOGICAL_STATUS_CODE"`
	MATCH                         float64 `json:"MATCH"`
	MATCHDESCRIPTION              string  `json:"MATCH_DESCRIPTION"`
	POSTALADDRESSCODE             string  `json:"POSTAL_ADDRESS_CODE"`
	POSTALADDRESSCODEDESCRIPTION  string  `json:"POSTAL_ADDRESS_CODE_DESCRIPTION"`
	POSTCODE                      string  `json:"POSTCODE"`
	POSTTOWN                      string  `json:"POST_TOWN"`
	RPC                           string  `json:"RPC"`
	STATUS                        string  `json:"STATUS"`
	THOROUGHFARENAME              string  `json:"THOROUGHFARE_NAME"`
	TOPOGRAPHYLAYERTOID           string  `json:"TOPOGRAPHY_LAYER_TOID"`
	UPRN                          string  `json:"UPRN"`
	XCOORDINATE                   float64 `json:"X_COORDINATE"`
	YCOORDINATE                   float64 `json:"Y_COORDINATE"`
}

//UprnResponse is the response from /places/v1/addresses/uprn
type UprnResponse struct {
	Header struct {
		Dataset      string `json:"dataset"`
		Epoch        string `json:"epoch"`
		Format       string `json:"format"`
		Lr           string `json:"lr"`
		Maxresults   int    `json:"maxresults"`
		Offset       int    `json:"offset"`
		Query        string `json:"query"`
		Totalresults int    `json:"totalresults"`
		URI          string `json:"uri"`
	} `json:"header"`
	Results []struct {
		DPA Dpa `json:"DPA"`
	} `json:"results"`
}

func FetchDpaInfo(uprn string, APIKey string) (Dpa, error) {
	if len(uprn) <= 0 {
		return Dpa{}, fmt.Errorf("uprn cannot be blank")
	}

	url := fmt.Sprintf("https://api.ordnancesurvey.co.uk/places/v1/addresses/uprn?uprn=%s&key=%s", uprn, APIKey)
	resp, err := http.Get(url)
	if err != nil {
		return Dpa{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Dpa{}, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return Dpa{}, errors.New(ErrorNotAuthorised)
	} else if resp.StatusCode != http.StatusOK {
		return Dpa{}, fmt.Errorf("response for UPRN %s was %s", uprn, resp.Status)
	}

	var jsonResponse UprnResponse
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return Dpa{}, err
	}

	if len(jsonResponse.Results) == 0 {
		return Dpa{}, errors.New(ErrorUprnNotFound)
	}

	return jsonResponse.Results[0].DPA, nil
}
