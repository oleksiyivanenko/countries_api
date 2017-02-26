package main

import (
  "fmt"
  "os"
  "strings"
  "encoding/json"
  "net/http"
  "io/ioutil"
)

type CountryName struct {
  Common string `json:"common"`
  Official string `json:"official"`
}

type Country struct {
  Name struct {
    Common string `json:"common"`
    Official string `json:"official"`
    Native map[string] CountryName `json:"native"`
  } `json:"name"`
  Tld []string `json:"tld"`

  Cca2 string `json:"cca2"`
  Ccn3 string `json:"ccn3"`
  Cca3 string `json:"cca3"`
  Cioc string `json:"cioc"`
  Currency []string `json:"currency"`
  CallingCode []string `json:"callingCode"`
  Capital string `json:"capital"`
  AltSpellings []string `json:"altSpellings"`
  Region string `json:"region"`
  Subregion string `json:"subregion"`
  Languages map[string] string `json:"languages"`

  Translations map[string] CountryName `json:"translations"`
  Latlng []float32 `json:"latlng"`
  Denonym string `json:"denonym"`
  Landlocked bool `json:"landlocked"`
  Borders []string `json:"borders"`
  Area float32 `json:"area"`
}

var (
  js []byte
  err error
)

var cca2Countries map[string] []byte
var cca3Countries map[string] []byte
var nameCountries map[string] []byte


func main() {
  filePath := "./countries.json"
  fmt.Printf("Reading file %s\n", filePath)

  js, err = ioutil.ReadFile(filePath)

  if err != nil {
    fmt.Printf("Error while reading file %s\n", filePath)
    os.Exit(1)
  }

  var countries []Country

  err2 := json.Unmarshal(js, &countries)
  if err2 != nil {
    fmt.Println("error:", err2)
    os.Exit(1)
  }

  cca2Countries = make(map[string] []byte)
  cca3Countries = make(map[string] []byte)
  nameCountries = make(map[string] []byte)

  for k := range countries {
    marshaledCountry, _ := json.Marshal(countries[k])

    cca2Countries[strings.ToLower(countries[k].Cca2)] = marshaledCountry
    cca3Countries[strings.ToLower(countries[k].Cca3)] = marshaledCountry
    nameCountries[strings.ToLower(countries[k].Name.Common)] = marshaledCountry
  }

  fmt.Printf("Serving\n")
  http.HandleFunc("/api/v1/countries", allCountries)
  http.HandleFunc("/api/v1/countries/", oneCountry)
  http.ListenAndServe(":8080", nil)
}

func allCountries(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func oneCountry(w http.ResponseWriter, r *http.Request) {
  index := r.URL.Path[len("/api/v1/countries/"):]
  index = strings.ToLower(index)
  w.Header().Set("Content-Type", "application/json")

  if len(index) == 2 {
    w.Write(cca2Countries[index])
  } else if len(index) == 3 {
    w.Write(cca3Countries[index])
  } else {
    w.Write(nameCountries[index])
  }
}
