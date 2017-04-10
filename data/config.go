package data

import (
	"os"
	"log"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// SQL Configuration
var sqlConf cloudsql
func init() {
	sqlConf.getConf()
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

type cloudsql struct {
	Connection	string	`yaml:"instance"`
	UserName	string 	`yaml:"user"`
	Password	string	`yaml:"paswd"`
}

func (cl cloudsql) String() string {
	return fmt.Sprintf("Google Cloud SQL Config:\n  conn:%s, user:%s, pass:%s",
		cl.Connection, cl.UserName, cl.Password)
}

func (c *cloudsql) getConf() *cloudsql {
	yamlFile, err := ioutil.ReadFile("./../config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}