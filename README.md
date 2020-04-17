# go-covid

## Description

Go Package to get information regarding the novel corona virus provided by Johns Hopkins university

Full Documentation can be found [here](https://ahmednafies.github.io/go-covid/)

## How to install
    go get -u github.com/ahmednafies/go-covid/covid

## How to use

### Get All Data


    data, err := covid.GetData()
    if err != nil {
        // Handle error here
    }
    fmt.Println(data)


### Get Status By Country Name

    data, err := covid.GetCountryByName("italy")
    if err != nil {
        // Handle error here
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

    list, err := covid.ListCountries()
    if err != nil {
        // Handle error here
    }

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

    data, err := covid.GetCountryById(113)

### Get Total Active Cases

    active, err := covid.GetTotalActive()

### Get Total Confirmed Cases

    confirmed, err := covid.GetTotalConfirmed()

### Get Total Recovered Cases

    recovered, err := covid.GetTotalRecovered()

### Get Total Deaths

    deaths, err := covid.GetTotalDeaths()
