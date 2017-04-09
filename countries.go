package main

import (
  "fmt"
  "os"
  "strings"
  "encoding/json"
  "net/http"
  "io/ioutil"

  "github.com/gorilla/mux"
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
var countriesShorltlist []byte


func main() {
  filePath := "./countries.json"
  fmt.Printf("Reading file %s\n", filePath)

  js, err = ioutil.ReadFile(filePath)

  if err != nil {
    fmt.Printf("Error while reading file %s\n", filePath)
    os.Exit(1)
  }

  var countries []Country

  err = json.Unmarshal(js, &countries)
  if err != nil {
    fmt.Println("error:", err)
    os.Exit(1)
  }

  cca2Countries = make(map[string] []byte)
  cca3Countries = make(map[string] []byte)
  nameCountries = make(map[string] []byte)

  var countries_short [][]string

  for k := range countries {
    marshaledCountry, _ := json.Marshal(countries[k])

    cca2Countries[strings.ToLower(countries[k].Cca2)] = marshaledCountry
    cca3Countries[strings.ToLower(countries[k].Cca3)] = marshaledCountry
    nameCountries[strings.ToLower(countries[k].Name.Common)] = marshaledCountry

    countries_short = append(countries_short, []string{countries[k].Cca2, countries[k].Name.Common})
  }

  countriesShorltlist, err = json.Marshal(countries_short)

  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/v1/countries", allCountries)
  router.HandleFunc("/v1/countries/{country_index}", oneCountry)

  http.Handle("/", router)
  http.ListenAndServe(":8080", nil)
}

func allCountries(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  _, full := r.URL.Query()["full"]
  if full {
    w.Write(js)
  } else {
    w.Write(countriesShorltlist)
  }
}

func oneCountry(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  country_index := vars["country_index"]
  country_index = strings.ToLower(country_index)
  w.Header().Set("Content-Type", "application/json")

  var response []byte
  var ok bool
  if len(country_index) == 2 {
    response, ok = cca2Countries[country_index]
  } else if len(country_index) == 3 {
    response, ok = cca3Countries[country_index]
  } else {
    response, ok = nameCountries[country_index]
  }

  if !ok {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("{\"error\": \"Not Found\"}"))
    return
  }

  w.Write(response)
}
