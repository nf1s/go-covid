package covid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _int int
var _string string
var _float float64

func TestGetCountryByName(t *testing.T) {

	data, err := GetCountryByName("italy")

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

	data, err := GetCountryById(50)
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

	all_data, err := GetData()
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

func TestListCountries(t *testing.T) {
	data, err := ListCountries()
	if err != nil {
		fmt.Println(err)
	}
	country := data[0]
	assert.IsType(t, _int, country.Attrs.Id)
	assert.IsType(t, _string, country.Attrs.Name)
}

func TestGetTotalActive(t *testing.T) {
	result, err := GetTotalActive()
	if err != nil {
		fmt.Println(err)
	}

	assert.IsType(t, _int, result)

}
func TestGetTotalConfirmed(t *testing.T) {
	result, err := GetTotalConfirmed()
	if err != nil {
		fmt.Println(err)
	}

	assert.IsType(t, _int, result)

}

func TestGetTotalRecovered(t *testing.T) {
	result, err := GetTotalRecovered()
	if err != nil {
		fmt.Println(err)
	}

	assert.IsType(t, _int, result)

}

func TestGetTotalDeaths(t *testing.T) {
	result, err := GetTotalDeaths()
	if err != nil {
		fmt.Println(err)
	}

	assert.IsType(t, _int, result)

}
