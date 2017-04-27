package data

import "log"

// In-memory caches
var cityRepo []City

// Load the data from DB to cache
func RefreshCity()  error {
  err := fetchCity(&cityRepo)
  if err == nil {
    log.Println("Cities cache refreshed")
  } else {
    log.Printf("Error refreshing cache: %v",err)
  }
  return err
}

func GetCityRepo() *[]City {
  return &cityRepo
}

func (city *[]City) Push(c City) {

}
