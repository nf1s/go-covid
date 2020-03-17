# go-covid

## Description

Go Package to get information regarding the novel corona virus provided by Johns Hopkins university

Full Documentation can be found [here](https://ahmednafies.github.io/go-covid/)

## How to install

    go get -u github.com/ahmednafies/go-covid/covid

## How to use

### Get All Data

    data := covid.GetData()
    fmt.Println(data)

### Get Status By Country Name

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

#### Result

    113
    Italy
    1584216796000
    21157
    1441
    17750
    1966
    41.8719

### List Countries

    covid.ListCountries()

#### Result

    [
        Attrs{
            Id: 40
            Name: Hungary
        },
        Attrs{
            Id: 113
            Name: Italy
        }
        ...
    ]

### Get Status by Country ID

    data := covid.GetCountryById(113)

### Get Total Active Cases

    active := covid.GetTotalActive()

### Get Total Confirmed Cases

    confirmed := covid.GetTotalConfirmed()

### Get Total Recovered Cases

    recovered := covid.GetTotalRecovered()

### Get Total Deaths

    deaths := covid.GetTotalDeaths()
