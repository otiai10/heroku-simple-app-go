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
			"LD_INCLUDE_PATH": os.Getenv("LD_INCLUDE_PATH"),
		})
	})

	http.ListenAndServe(":"+port, router)
}
