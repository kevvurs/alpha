package service

import (
	"os"
	"log"
	"database/sql"
	"fmt"
	"net/http"
	"bytes"
)

const USE_DATABASE string = "USE CC_LOCALS"
const SELECT_PLACES = "SELECT * FROM places"

// Cache
func init() {

}

//
func Cache() {
	connectionName := sysVar("CLOUDSQL_CONNECTION_NAME")
	user := sysVar("CLOUDSQL_USER")
	password := sysVar("CLOUDSQL_PASSWORD")

	// Connect to Google Cloud SQL
	log.Printf("Connecting to %s ", connectionName)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/", user, password, connectionName))
	if err != nil {
		log.Panicf("Could not open db: %v", err)
		return
	}
	defer db.Close()

	// Initialize local variables and namespace
	rows, err := db.Query(USE_DATABASE)
	if err != nil {
		log.Panicf("Could not query db\n  query:%s\n  error:%v", USE_DATABASE, err)
		return
	}
	defer rows.Close()

	rows, err = db.Query(SELECT_PLACES)
	if err != nil {
		log.Panicf("Could not query db\n  query:%s\n  error:%v", SELECT_PLACES, err)
		return
	}

	buf := bytes.NewBufferString("Databases:\n")
	for rows.Next() {
		var city, country, description string
		var score int32
		var pop int64
		if err := rows.Scan(&city, &country, &description, &score, &pop); err != nil {
			http.Error(w, fmt.Sprintf("Could not scan result: %v", err), 500)
			return
		}
		fmt.Fprintf(buf, "- %s\n", dbName)
	}
	w.Write(buf.Bytes())
}

func sysVar(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}