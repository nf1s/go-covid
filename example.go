package main

import (
	"fmt"
	"go-covid/covid"
)

func main() {
	data, err := covid.GetCountryByName("italy")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data.Attrs.Id)
	fmt.Println(data.Attrs.Country)
	fmt.Println(data.Attrs.LastUpdate)
	fmt.Println(data.Attrs.Confirmed)
	fmt.Println(data.Attrs.Deaths)
	fmt.Println(data.Attrs.Active)
	fmt.Println(data.Attrs.Recovered)
	fmt.Println(data.Attrs.Latitude)
	fmt.Println(data.Attrs.Longitude)

	active, err := covid.GetTotalActive()
	if err != nil {
		fmt.Println(err)
	}
	confirmed, err := covid.GetTotalConfirmed()
	if err != nil {
		fmt.Println(err)
	}
	recovered, err := covid.GetTotalRecovered()
	if err != nil {
		fmt.Println(err)
	}
	deaths, err := covid.GetTotalDeaths()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(active)
	fmt.Println(confirmed)
	fmt.Println(recovered)
	fmt.Println(deaths)

}
