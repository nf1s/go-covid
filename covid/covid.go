package covid

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var BaseUrl = "https://services1.arcgis.com"

var Path = "/0MSEUqKaxRlEPj5g/arcgis/rest/services/ncov_cases/FeatureServer/2/query"

type Response struct {
	Values []Case `json:"features"`
}
type Case struct {
	Attrs struct {
		Id         int     `json:"OBJECTID"`
		Country    string  `json:"Country_Region"`
		LastUpdate int     `json:"Last_Update"`
		Latitude   float64 `json:"Lat"`
		Longitude  float64 `json:"Long_"`
		Confirmed  int     `json:"Confirmed"`
		Active     int     `json:"Active"`
		Recovered  int     `json:"Recovered"`
		Deaths     int     `json:"Deaths"`
	} `json:"attributes"`
}

type CountryResponse struct {
	Values []Country `json:"features"`
}

type Country struct {
	Attrs struct {
		Id   int    `json:"OBJECTID"`
		Name string `json:"Country_Region"`
	} `json:"attributes"`
}
type StatsResponse struct {
	Values []Stats `json:"features"`
}

type Stats struct {
	Attrs struct {
		Value int `json:"value"`
	} `json:"attributes"`
}

func getAllUrl() string {
	_url, _ := url.Parse(BaseUrl)
	_url.Path += Path
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

func getCountryUrl(id int) string {
	_url, _ := url.Parse(BaseUrl)
	_url.Path += Path
	parameters := url.Values{}
	parameters.Add("f", "json")
	parameters.Add("where", "OBJECTID ="+strconv.Itoa(id))
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

func getTotalUrl(field string) string {
	_url, _ := url.Parse(BaseUrl)
	_url.Path += Path
	parameters := url.Values{}
	parameters.Add("f", "json")
	parameters.Add("where", "Confirmed > 0")
	parameters.Add("returnGeometry", "false")
	parameters.Add("spatialRel", "esriSpatialRelIntersects")
	parameters.Add("outFields", "*")
	parameters.Add("outStatistics", fmt.Sprintf("[{\"statisticType\":\"sum\",\"onStatisticField\":\"%s\",\"outStatisticFieldName\":\"value\"}]", field))
	parameters.Add("cacheHint", "true")
	_url.RawQuery = parameters.Encode()
	return _url.String()
}

func GetData() []Case {
	resp, err := http.Get(getAllUrl())
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var res Response
	json.Unmarshal(body, &res)
	if err != nil {
		log.Fatalln(err)
	}

	return res.Values
}

func GetCountryById(id int) Case {
	resp, err := http.Get(getCountryUrl(id))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var res Response
	json.Unmarshal(body, &res)
	if err != nil {
		log.Fatalln(err)
	}

	return res.Values[0]
}

func ListCountries() []Country {
	resp, err := http.Get(getAllUrl())
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var res CountryResponse
	json.Unmarshal(body, &res)
	if err != nil {
		log.Fatalln(err)
	}

	return res.Values
}

func getCountryId(name string) (int, error) {
	countries := ListCountries()
	for _, country := range countries {
		if strings.ToLower(country.Attrs.Name) == strings.ToLower(name) {
			return country.Attrs.Id, nil
		}
	}
	return -1, errors.New("Wrong country name")
}

func GetCountryByName(name string) Case {
	id, err := getCountryId(name)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Get(getCountryUrl(id))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var res Response
	json.Unmarshal(body, &res)
	if err != nil {
		log.Fatalln(err)
	}

	return res.Values[0]
}

func getTotal(field string) Stats {
	resp, err := http.Get(getTotalUrl(field))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var res StatsResponse
	json.Unmarshal(body, &res)
	if err != nil {
		log.Fatalln(err)
	}

	return res.Values[0]
}

func GetTotalActive() int {
	data := getTotal("Active")
	return data.Attrs.Value
}

func GetTotalConfirmed() int {
	data := getTotal("Confirmed")
	return data.Attrs.Value
}
func GetTotalRecovered() int {
	data := getTotal("Recovered")
	return data.Attrs.Value
}
func GetTotalDeaths() int {
	data := getTotal("Deaths")
	return data.Attrs.Value
}
