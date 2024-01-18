package main

import (
	"fmt"
	"log"
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

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
