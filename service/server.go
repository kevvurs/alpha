package service

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"log"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	log.Println("Launching Go server")

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	// repo := initRepository(appEnv)

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

// Routing with Gorilla Mux
func initRoutes(mx *mux.Router, formatter *render.Render) {
	log.Println("Routing API endpoints")
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/publication", getData(formatter)).Methods("GET")
	mx.HandleFunc("/publication/{pubid}", getData(formatter)).Methods("GET")
	// mx.HandleFunc("/matches", createMatchHandler(formatter, repo)).Methods("POST")
	// mx.HandleFunc("/matches", getMatchListHandler(formatter, repo)).Methods("GET")
	// mx.HandleFunc("/matches/{id}", getMatchDetailsHandler(formatter, repo)).Methods("GET")
	// mx.HandleFunc("/matches/{id}/moves", addMoveHandler(formatter, repo)).Methods("POST")
}
