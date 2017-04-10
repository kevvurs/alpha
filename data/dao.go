package data

import (
	"log"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Get all cities in db
func FetchCity() ([]City, error){
	var cityList []City

	// Connect to Google Cloud SQL
	log.Printf("Connecting to %s ", connectionName)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/", user, password, connectionName))
	if err != nil {
		log.Panicf("Could not open db: %v", err)
		return cityList, err
	}
	defer db.Close()

	// Initialize local variables and namespace
	rows, err := db.Query(use_database)
	if err != nil {
		log.Panicf("Could not query db\n  query:%s\n  error:%v", use_database, err)
		return cityList, err
	}
	defer rows.Close()

	rows, err = db.Query(select_places)
	if err != nil {
		log.Panicf("Could not query db\n  query:%s\n  error:%v", select_places, err)
		return cityList, err
	}

	// Marshalling to structure
	for rows.Next() {
		var name, country, description string
		var score int
		var pop int64
		if err := rows.Scan(&name, &country, &description, &score, &pop); err != nil {
			log.Panicf("Could not scan result: %v", err)
			return cityList, err
		}

		// Collect additional rows
		city := City{
			Name: name,
			Country: country,
			Description: description,
			Score: score,
			Pop: pop,
		}
		cityList = append(cityList, city)
	}

	return cityList, nil
}