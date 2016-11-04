package main

import (
	"fmt"
	"net/http"
	"os"

	m "github.com/otiai10/marmoset"
)

func main() {

	port := os.Getenv("PORT")

	router := m.NewRouter()
	router.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	http.ListenAndServe(":"+port, router)
}
