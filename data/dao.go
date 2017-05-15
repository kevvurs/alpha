package data

import (
	"database/sql"
	proxy "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"log"
)

// Fetch all data and update cache references
// Publications no longer in DB will not be removed
func fetchAll(pubs *PubCache) error {
	// SQL pointers
	var db *sql.DB
	var tx *sql.Tx
	var rows *sql.Rows
	var err error

	// Connect to Google Cloud SQL
	log.Printf("Connecting to %s ", sqlConf.Connection)
	if db, err = proxy.DialPassword(sqlConf.Connection, sqlConf.UserName, sqlConf.Password); err != nil {
		log.Panicf("Could not open db: error:%v\n", err)
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
	if rows, err = tx.Query(sqlConf.Stmnt.Use_database); err != nil {
		log.Printf("Query failed: %s\n", sqlConf.Stmnt.Use_database)
		tx.Rollback()
		return err
	}
	defer rows.Close()

	if rows, err = tx.Query(sqlConf.Stmnt.Select_publications); err != nil {
		log.Printf("Query failed: %s\n", sqlConf.Stmnt.Select_publications)
		tx.Rollback()
		return err
	}

	// Marshalling to structure
	for rows.Next() {
		var publisher, home, imgref, owner string
		var hits, ycred, ncred, pubId int
		var quality float32
		if err := rows.Scan(&publisher, &home, &imgref, &hits, &quality, &ycred, &ncred, &owner, &pubId); err != nil {
			log.Panicf("Error: cannot read from resultset: %v\n", err)
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
			Exists:    true,
		}
		pubs.cache[pubId] = pub
	}

	// Closeout transaction
	if err = tx.Commit(); err != nil {
		log.Printf("Warning: SQL transaction not properly commited: %v\n", err)
		return err
	}
	return nil
}

// Persist data to DB, with clobbering
func upsert(pub *Publication) error {
	db, tx, rows, err := openConnection()
	defer closeConnection(db, tx, rows, err)
	var stmt *sql.Stmt

	// Prepare the SQL statement
	if stmt, err = tx.Prepare(sqlConf.Stmnt.Insert_update); err != nil {
		log.Println("Upsertion failed to prepare: %v\n", err)
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(pub.Publisher, pub.Home, pub.Imgref, pub.Hits,
		pub.Quality, pub.Ycred, pub.Ncred, pub.Owner, pub.PubId,
		pub.Hits, pub.Ycred, pub.Ncred, pub.Quality); err != nil {
		log.Println("Upsertion failed to execute: %v\n", err)
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Println("Upsertion failed to commit: %v\n", err)
		return err
	}

	return err
}

// Persist data to DB, with clobbering
func delete(id *int) error {
	db, tx, rows, err := openConnection()
	defer closeConnection(db, tx, rows, err)
	var stmt *sql.Stmt

	// Prepare the SQL statement
	if stmt, err = tx.Prepare(sqlConf.Stmnt.Delete_publication); err != nil {
		log.Println("Deletion failed to prepare: %v\n", err)
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		log.Println("Deletion failed to execute: %v\n", err)
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Println("Deletion failed to commit: %v\n", err)
		return err
	}

	return err
}

// Opens all the crucial streams for DB access
func openConnection() (db *sql.DB, tx *sql.Tx, rows *sql.Rows, err error) {
	// Connect to Google Cloud SQL
	log.Printf("Connecting to %s ", sqlConf.Connection)
	if db, err = proxy.DialPassword(sqlConf.Connection, sqlConf.UserName, sqlConf.Password); err != nil {
		log.Panicf("Could not open db: error:%v\n", err)
		return
	}

	tx, err = db.Begin()
	if err != nil {
		log.Println("Could not open transaction stream")
		return
	}

	// Initialize local variables and namespace
	if rows, err = tx.Query(sqlConf.Stmnt.Use_database); err != nil {
		log.Printf("Query failed: %s\n", sqlConf.Stmnt.Use_database)
		// tx.Rollback()
		return
	}

	return
}

// Cleans up connection artifacts
// Audit the transaction rollback
func closeConnection(db *sql.DB, tx *sql.Tx, rows *sql.Rows, err error) {
	if err := tx.Rollback(); err != nil {
		if err != sql.ErrTxDone {
			log.Println("Potential errors while rolling back SQL transaction")
		}
	}

	rows.Close()
	db.Close()
}
