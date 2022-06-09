package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	// reader := bufio.NewReader(os.Stdin)

	// for {
	// 	println("What is your name")
	// 	text, _ := reader.ReadString('\n')
	// 	println("Hello ", text)
	// }

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/users", usersPage)

	port := ":8080" //на каком порту запуститься сервер
	println("server listem on port", port)
	err := http.ListenAndServe(port, nil) // запускаем цикл
	if err != nil {
		log.Fatal("ListenAndServe ", err)
	}

}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsFired   bool
}

func mainPage(w http.ResponseWriter, h *http.Request) {
	// user := User{"Anton", "TT"}
	// js, _ := json.Marshal(user)

	//пример рендеринга с файла
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
	}
}

func usersPage(w http.ResponseWriter, h *http.Request) {
	users := []User{User{"Anton", "Sidorov", false}, User{"Kolya", "Vasechkin", true}}
	tmpl, err := template.ParseFiles("static/users.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := tmpl.Execute(w, users); err != nil {
		http.Error(w, err.Error(), 400)
	}
}
