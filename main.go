package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := ":8080"

	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hihihihi")
	})

	http.ListenAndServe(port, nil)
}
