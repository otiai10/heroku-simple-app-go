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
			"message": "Hello!",
			"FOO":     os.Getenv("FOO"),
		})
	})

	http.ListenAndServe(":"+port, router)
}
