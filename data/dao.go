package data

import (
	"database/sql"
	proxy "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"log"
)

// Get all cities in db
func fetchCity(pubs *PubCache) error {
	pubs.mux.Lock()
	defer pubs.mux.Unlock()

	// SQL pointers
	var db *sql.DB
	var tx *sql.Tx
	var rows *sql.Rows
	var err error

	// Connect to Google Cloud SQL
	log.Printf("Connecting to %s ", sqlConf.Connection)
	if db, err = proxy.DialPassword(sqlConf.Connection, sqlConf.UserName, sqlConf.Password); err != nil {
		log.Panicf("Could not open db: error:%v", err)
		return err
	}
	defer db.Close()
	log.Println("Connected")

	tx, err = db.Begin()
	if err != nil {
		log.Println("Could not open transaction stream")
		return err
	}

	// Initialize local variables and namespace
	if rows, err = tx.Query(use_database); err != nil {
		log.Printf("Query failed: %s\n", use_database)
		tx.Rollback()
		return err
	}
	defer rows.Close()

	if rows, err = tx.Query(select_places); err != nil {
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
		pub := Publication{
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
		pubs.cache[pubId] = pub
	}

	// Closeout transaction
	if err = tx.Commit(); err != nil {
		log.Printf("Warning: SQL transaction not properly commited: %v", err)
		return err
	}
	return nil
}
