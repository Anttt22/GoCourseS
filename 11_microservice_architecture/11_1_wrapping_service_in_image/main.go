package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "hello")
		if err != nil {
			fmt.Fprintf(w, "errorHandleFunc: #{err.Error()}")
		}

	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("failed to start server: #{err.Error()}")
	}
}
