package main

import (
	"net/http"
	"os"

	m "github.com/otiai10/marmoset"
)

func main() {

	port := os.Getenv("PORT")

	router := m.NewRouter()
	router.GET("/", func(w http.ResponseWriter, req *http.Request) {
		m.RenderJSON(w, http.StatusOK, m.P{
			"TESSDATA_PREFIX": os.Getenv("TESSDATA_PREFIX"),
			"INCLUDE_PATH":    os.Getenv("INCLUDE_PATH"),
			"LIBRARY_PATH":    os.Getenv("LIBRARY_PATH"),
			"LD_LIBRARY_PATH": os.Getenv("LD_LIBRARY_PATH"),
			"HOME":            os.Getenv("HOME"),
		})
	})

	http.ListenAndServe(":"+port, router)
}
