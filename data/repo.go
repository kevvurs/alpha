package data

import "log"

// In-memory caches
var cityRepo []Publication

// Load the data from DB to cache
func RefreshCity()  error {
  err := fetchCity(&cityRepo)
  if err == nil {
    log.Println("Pub cache refreshed")
  } else {
    log.Printf("Error: refreshing cache: %v",err)
  }
  return err
}

func GetCityRepo() *[]Publication {
  return &cityRepo
}

// func (city *[]City) Push(c City)
