package main

import (
	"fmt"
	"go-covid/covid"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _int int
var _string string
var _float float64

func TestGetCountryByName(t *testing.T) {

	data, err := covid.GetCountryByName("italy")

	if err != nil {
		fmt.Println(err)
	}

	assert.IsType(t, _int, data.Attrs.Id)
	assert.IsType(t, _string, data.Attrs.Country)
	assert.IsType(t, _int, data.Attrs.LastUpdate)
	assert.IsType(t, _int, data.Attrs.Confirmed)
	assert.IsType(t, _int, data.Attrs.Deaths)
	assert.IsType(t, _int, data.Attrs.Active)

	assert.IsType(t, _int, data.Attrs.Recovered)
	assert.IsType(t, _float, data.Attrs.Latitude)
	assert.IsType(t, _float, data.Attrs.Longitude)
	assert.Equal(t, "Italy", data.Attrs.Country)

}

func TestGetCountryById(t *testing.T) {

	data, err := covid.GetCountryById(50)
	if err != nil {
		fmt.Println(err)
	}

	assert.IsType(t, _int, data.Attrs.Id)
	assert.IsType(t, _string, data.Attrs.Country)
	assert.IsType(t, _int, data.Attrs.LastUpdate)
	assert.IsType(t, _int, data.Attrs.Confirmed)
	assert.IsType(t, _int, data.Attrs.Deaths)
	assert.IsType(t, _int, data.Attrs.Active)
	assert.IsType(t, _int, data.Attrs.Recovered)
	assert.IsType(t, _float, data.Attrs.Latitude)
	assert.IsType(t, _float, data.Attrs.Longitude)

}

func TestGetAll(t *testing.T) {

	all_data, err := covid.GetData()
	if err != nil {
		fmt.Println(err)
	}
	data := all_data[0]
	assert.IsType(t, _int, data.Attrs.Id)
	assert.IsType(t, _string, data.Attrs.Country)
	assert.IsType(t, _int, data.Attrs.LastUpdate)
	assert.IsType(t, _int, data.Attrs.Confirmed)
	assert.IsType(t, _int, data.Attrs.Deaths)
	assert.IsType(t, _int, data.Attrs.Active)
	assert.IsType(t, _int, data.Attrs.Recovered)
	assert.IsType(t, _float, data.Attrs.Latitude)
	assert.IsType(t, _float, data.Attrs.Longitude)

}

func ListCountries(t *testing.T) {
	data, err := covid.ListCountries()
	if err != nil {
		fmt.Println(err)
	}
	country := data[0]
	assert.IsType(t, _int, country.Attrs.Id)
	assert.IsType(t, _string, country.Attrs.Name)
}

func TestGetTotalActive(t *testing.T) {
	result, err := covid.GetTotalActive()
	if err != nil {
		fmt.Println(err)
	}

	assert.IsType(t, _int, result)

}
func TestGetTotalConfirmed(t *testing.T) {
	result, err := covid.GetTotalConfirmed()
	if err != nil {
		fmt.Println(err)
	}

	assert.IsType(t, _int, result)

}

func TestGetTotalRecovered(t *testing.T) {
	result, err := covid.GetTotalRecovered()
	if err != nil {
		fmt.Println(err)
	}

	assert.IsType(t, _int, result)

}

func TestGetTotalDeaths(t *testing.T) {
	result, err := covid.GetTotalDeaths()
	if err != nil {
		fmt.Println(err)
	}

	assert.IsType(t, _int, result)

}
