package main

import (
  "fmt"
  "os"
  "encoding/json"
  "net/http"
  "io/ioutil"
)

type Country struct {
  Name struct {
    Common string `json:"common"`
    Official string `json:"official"`
    // Native map[string] map[string] string
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
var countriesMap map[string] []byte


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

  countriesMap = make(map[string] []byte)
  for k := range countries {
    countriesMap[countries[k].Cca2], _ = json.Marshal(countries[k])
  }

  fmt.Printf("Serving\n")
  http.HandleFunc("/api/countries", allCountries)
  http.HandleFunc("/api/countries/", oneCountry)
  http.ListenAndServe(":8080", nil)
}

func allCountries(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func oneCountry(w http.ResponseWriter, r *http.Request) {
  index := r.URL.Path[len("/api/countries/"):]
  w.Header().Set("Content-Type", "application/json")
  w.Write(countriesMap[index])
}
