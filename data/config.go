package data

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

// SQL Configuration
var sqlConf = initConf()

func sysVar(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}

func initConf() cloudsql {
	c := new(cloudsql)
	yamlFile, err := ioutil.ReadFile("./../config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return *c
}
