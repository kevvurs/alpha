package data

import (
	"log"
	"database/sql"
	proxy "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
)

// Get all cities in db
func FetchCity() ([]City, error){
	// return vals
	var cityList []City
	var err error

	// SQL pointers
	var 	db 		*sql.DB
	var 	tx 		*sql.Tx
	var 	rows 	*sql.Rows

	// Connect to Google Cloud SQL
	log.Printf("Connecting to %s ", sqlConf.Connection)
	db, err = proxy.DialPassword(sqlConf.Connection, sqlConf.UserName, sqlConf.Password)
	log.Println("Connected")
	if err != nil {
		log.Panicf("Could not open db: error:%v", err)
		return cityList, err
	}
	defer db.Close()

	tx, err = db.Begin()
	if err != nil {
		log.Panicf("Could not open transaction stream: error:%v", err)
		return cityList, err
	}

	// Initialize local variables and namespace
	rows, err = tx.Query(use_database)
	if err != nil {
		log.Panicf("Could not query db\n  query:%s\n  error:%v", use_database, err)
		return cityList, err
	}
	defer rows.Close()

	rows, err = tx.Query(select_places)
	if err != nil {
		log.Panicf("Could not query db\n  query:%s\n  error:%v", select_places, err)
		return cityList, err
	}

	// Marshalling to structure
	for rows.Next() {
		var name, country, description, timezone string
		var score int
		var pop int64
		if err := rows.Scan(&name, &country, &description, &score, &timezone, &pop); err != nil {
			log.Panicf("Could not scan result: %v", err)
			return cityList, err
		}

		// Collect additional rows
		city := City{
			Name: name,
			Country: country,
			Description: description,
			Score: score,
			Timezone: timezone,
			Pop: pop,
		}
		cityList = append(cityList, city)
	}
	return cityList, nil
}


/* Appenngine heap pile
	"fmt"
	 _ "github.com/go-sql-driver/mysql"
	 db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/", sqlConf.UserName, sqlConf.Password, sqlConf.Connection))
*/
