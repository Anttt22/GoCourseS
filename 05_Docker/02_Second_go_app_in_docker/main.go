package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/docker", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "hello from docker 2")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("failed to start server: #{err.Errr()}")
	}
}
