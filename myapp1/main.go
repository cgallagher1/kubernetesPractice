package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, os.Getenv("MESSAGE"))
	})

	port := fmt.Sprintf(":%s", os.Getenv("LISTENANDSERVE"))
	http.ListenAndServe(port, nil)
}
