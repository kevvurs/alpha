package data

import (
	"os"
	"log"
)

// SQL Configuration
var connectionName, user, password string
func init() {
	connectionName = sysVar("CLOUDSQL_CONNECTION_NAME")
	user = sysVar("CLOUDSQL_USER")
	password = sysVar("CLOUDSQL_PASSWORD")
}

// Query constants
const (
	use_database string = "USE cc_locals"
	select_places = "SELECT * FROM places"
)
func sysVar(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}