package main

import (
	"fmt"
	"net/http"

	m "github.com/otiai10/marmoset"
)

func main() {
	router := m.NewRouter()
	router.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	http.ListenAndServe(":5000", router)
}
