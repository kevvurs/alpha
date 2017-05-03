package service

import (
	"github.com/unrolled/render"
	"net/http"

	data "github.com/kevvurs/alpha/data"
)

// Ping status for Go server
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Ping string }{"OK"})
	}
}

func locationListHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		city, err := data.FetchCity()
		if err == nil {
			formatter.JSON(w, http.StatusOK, city)
		} else {
			formatter.JSON(w, http.StatusNotFound, err.Error())
		}
	}
}
