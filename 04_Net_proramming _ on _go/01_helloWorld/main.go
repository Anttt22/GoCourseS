package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	Listen = "0.0.0.0:8080"
)

func main() {
	http.HandleFunc("/", handleIndex) //принимает адресс и фуекцию handler
	http.HandleFunc("/time", handleTime)

	fmt.Printf("listen on %s\n", Listen) //выводим где слушает сервер
	http.ListenAndServe(Listen, nil)     // запускаем луп где будем слушать клиента
	fmt.Println("end")
}

//получает текущее время
func handleTime(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintf(w, "Time %s", t.Format(time.RFC3339Nano))
}

//
func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

// README cmd curl localhost:8080
