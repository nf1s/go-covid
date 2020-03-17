package main

import (
	"fmt"
	"go-covid/covid"
)

func main() {
	data := covid.GetCountryByName("italy")
	fmt.Println(data.Attrs.Id)
	fmt.Println(data.Attrs.Country)
	fmt.Println(data.Attrs.LastUpdate)
	fmt.Println(data.Attrs.Confirmed)
	fmt.Println(data.Attrs.Deaths)
	fmt.Println(data.Attrs.Active)
	fmt.Println(data.Attrs.Recovered)
	fmt.Println(data.Attrs.Latitude)
	fmt.Println(data.Attrs.Longitude)

	active := covid.GetTotalActive()
	confirmed := covid.GetTotalConfirmed()
	recovered := covid.GetTotalRecovered()
	deaths := covid.GetTotalDeaths()

	fmt.Println(active)
	fmt.Println(confirmed)
	fmt.Println(recovered)
	fmt.Println(deaths)

}
