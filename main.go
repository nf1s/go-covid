package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"net/url"

)



func get_all_url() string{
	_url, _ := url.Parse("https://services1.arcgis.com")
	_url.Path += "/0MSEUqKaxRlEPj5g/arcgis/rest/services/ncov_cases/FeatureServer/2/query"

	parameters := url.Values{}
    parameters.Add("f", "json")
    parameters.Add("where", "Confirmed > 0")
    parameters.Add("returnGeometry", "false")
    parameters.Add("spatialRel", "esriSpatialRelIntersects")
    parameters.Add("outFields", "*")
    parameters.Add("orderByFields", "Confirmed desc")
    parameters.Add("resultOffset", "0")
    parameters.Add("resultRecordCount", "200")
    parameters.Add("cacheHint", "true")

	_url.RawQuery = parameters.Encode()

	return _url.String()
}

func main() {
	MakeRequest()
}

func MakeRequest() {
	resp, err := http.Get(get_all_url())
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
