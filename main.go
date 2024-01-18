package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hihihihi")
	})

	http.ListenAndServe(port, nil)
}
