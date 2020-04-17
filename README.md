# go-covid

[![CircleCI](https://circleci.com/gh/ahmednafies/go-covid.svg?style=shield)](https://circleci.com/gh/ahmednafies/go-covid) [![codecov](https://codecov.io/gh/ahmednafies/go-covid/branch/master/graph/badge.svg)](https://codecov.io/gh/ahmednafies/go-covid)

## Description

Go Package to get information regarding the novel corona virus provided by Johns Hopkins university

Full Documentation can be found [here](https://ahmednafies.github.io/go-covid/)

## How to install

```go
go get -u github.com/ahmednafies/go-covid/covid
```

## How to use

### Get All Data

```go
data, err := covid.GetData()
if err != nil {
    // Handle error here
}
fmt.Println(data)
```

### Get Status By Country Name

```go
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
```

#### Result

```go
113
Italy
1584216796000
21157
1441
17750
1966
41.8719
```

### List Countries

```go
list, err := covid.ListCountries()
if err != nil {
    // Handle error here
}
```

#### Result

```go
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
```

### Get Status by Country ID

```go
data, err := covid.GetCountryById(113)
```

### Get Total Active Cases

```go
active, err := covid.GetTotalActive()
```

### Get Total Confirmed Cases

```go
confirmed, err := covid.GetTotalConfirmed()
```

### Get Total Recovered Cases

```go
recovered, err := covid.GetTotalRecovered()
```

### Get Total Deaths

```go
deaths, err := covid.GetTotalDeaths()
```
