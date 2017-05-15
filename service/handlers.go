package service

import (
	"strconv"
	"github.com/unrolled/render"
	"net/http"
	"github.com/kevvurs/alpha/data"
	"github.com/gorilla/mux"
)

// Ping status for Go server
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Ping string }{"OK"})
	}
}

func getData(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		v := vars["pubid"]

		if pubId, err := strconv.Atoi(v); err == nil {
			cache := data.GetRepo();
			if p := cache.Pull(&pubId); p.Exists {
				formatter.JSON(w, http.StatusOK, p)
			} else {
				formatter.JSON(w, http.StatusNotFound, error("Cannot retrieve data for ID"))
			}
		} else {
			formatter.JSON(w, http.StatusNotFound, error("Invalid path parameter"))
		}
	}
}
