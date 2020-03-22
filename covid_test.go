package main

import (
	"go-covid/covid"
	"testing"
	"github.com/stretchr/testify/assert"

)

var _int int
var _string string
var _float float64 

func TestGetCountryByName(t *testing.T) {

	data := covid.GetCountryByName("italy")

	assert.IsType(t,_int,data.Attrs.Id)
	assert.IsType(t,_string,data.Attrs.Country)
	assert.IsType(t,_int,data.Attrs.LastUpdate)
	assert.IsType(t,_int,data.Attrs.Confirmed)
	assert.IsType(t,_int,data.Attrs.Deaths)
	assert.IsType(t,_int,data.Attrs.Active)

	assert.IsType(t,_int,data.Attrs.Recovered)
	assert.IsType(t,_float,data.Attrs.Latitude)
	assert.IsType(t,_float,data.Attrs.Longitude)
	assert.Equal(t, "Italy", data.Attrs.Country)

}

func TestGetCountryById(t *testing.T) {


	data := covid.GetCountryById(50)
	assert.IsType(t,_int,data.Attrs.Id)
	assert.IsType(t,_string,data.Attrs.Country)
	assert.IsType(t,_int,data.Attrs.LastUpdate)
	assert.IsType(t,_int,data.Attrs.Confirmed)
	assert.IsType(t,_int,data.Attrs.Deaths)
	assert.IsType(t,_int,data.Attrs.Active)
	assert.IsType(t,_int,data.Attrs.Recovered)
	assert.IsType(t,_float,data.Attrs.Latitude)
	assert.IsType(t,_float,data.Attrs.Longitude)

}

func TestGetAll(t *testing.T) {

	all_data := covid.GetData()
	data := all_data[0]
	assert.IsType(t,_int,data.Attrs.Id)
	assert.IsType(t,_string,data.Attrs.Country)
	assert.IsType(t,_int,data.Attrs.LastUpdate)
	assert.IsType(t,_int,data.Attrs.Confirmed)
	assert.IsType(t,_int,data.Attrs.Deaths)
	assert.IsType(t,_int,data.Attrs.Active)
	assert.IsType(t,_int,data.Attrs.Recovered)
	assert.IsType(t,_float,data.Attrs.Latitude)
	assert.IsType(t,_float,data.Attrs.Longitude)

}

func ListCountries(t *testing.T) {
	data := covid.ListCountries()
	country := data[0]
	assert.IsType(t,_int,country.Attrs.Id)
	assert.IsType(t,_string,country.Attrs.Name)
}

func TestGetTotalActive(t *testing.T) {
	result := covid.GetTotalActive()
	assert.IsType(t,_int,result)

}
func TestGetTotalConfirmed(t *testing.T) {
	result := covid.GetTotalConfirmed()
	assert.IsType(t,_int,result)

}

func TestGetTotalRecovered(t *testing.T) {
	result := covid.GetTotalRecovered()
	assert.IsType(t,_int,result)

}

func TestGetTotalDeaths(t *testing.T) {
	result := covid.GetTotalDeaths()
	assert.IsType(t,_int,result)

}