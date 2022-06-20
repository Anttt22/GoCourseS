package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/docker", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello World from Docker container </h1>")
	})
	//выводит по адремсу /docker hello word form docker container
	// но чтобы это было из контаинера надо созлать образ(изображение)
	// создадим докер файл
	http.ListenAndServe(":8080", nil)
}
