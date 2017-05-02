package data

import (
	"database/sql"
	proxy "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"log"
)

// Get all cities in db
func fetchCity(cityList *[]Publication) error {
	var err error

	// SQL pointers
	var db *sql.DB
	var tx *sql.Tx
	var rows *sql.Rows
	//proxy.Init
	// Connect to Google Cloud SQL
	log.Printf("Connecting to %s ", sqlConf.Connection)
	db, err = proxy.DialPassword(sqlConf.Connection, sqlConf.UserName, sqlConf.Password)
	log.Println("Connected")
	if err != nil {
		log.Panicf("Could not open db: error:%v", err)
		return err
	}
	defer db.Close()

	tx, err = db.Begin()
	if err != nil {
		log.Println("Could not open transaction stream")
		return err
	}

	// Initialize local variables and namespace
	rows, err = tx.Query(use_database)
	if err != nil {
		log.Printf("Query failed: %s\n", use_database)
		tx.Rollback()
		return err
	}
	defer rows.Close()

	rows, err = tx.Query(select_places)
	if err != nil {
		log.Printf("Query failed: %s\n", select_places)
		tx.Rollback()
		return err
	}

	// Marshalling to structure
	for rows.Next() {
		var publisher, home, imgref, owner string
		var hits, ycred, ncred, pubId int
		var quality float32
		if err := rows.Scan(&publisher, &home, &imgref, &hits, &quality, &ycred, &ncred, &owner, &pubId); err != nil {
			log.Panicf("Error: cannot read from resultset: %v", err)
			tx.Rollback()
			return err
		}

		// Collect additional rows
		city := Publication{
			Publisher: publisher,
			Home:      home,
			Imgref:    imgref,
			Hits:      hits,
			Quality:   quality,
			Ycred:     ycred,
			Ncred:     ncred,
			Owner:     owner,
			PubId:     pubId,
		}
		*cityList = append(*cityList, city)
	}

	// Closeout transaction
	err = tx.Commit()
	if err != nil {
		log.Printf("Warning: SQL transaction not properly commited: %v", err)
		return err
	}

	return nil
}
