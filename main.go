package main


import (
	"os"
	"log"
	service "github.com/kevvurs/alpha/service"
)

func main() {
	log.Println("Running alpha from main")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	server := service.NewServer()
	server.Run(":" + port)
}
