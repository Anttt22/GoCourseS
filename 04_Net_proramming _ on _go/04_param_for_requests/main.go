package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// параметры из url  идут после ?   /user?name=kolya&id=22332

	// http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {

	// 	name := r.URL.Query().Get("name")
	// 	age := r.URL.Query().Get("id")
	// 	fmt.Fprintf(w, "Имя: %s Возраст: %s", name, age)
	// })

	//Через FORMDATA ФОРМУДАТА
	// параметры через форму получаемую из user.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "user.html")

	})

	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("username")
		age := r.FormValue("userage")

		fmt.Fprintf(w, "from form: name of user %s and age is %s", name, age)
	})
	//HTTP Body  тело запроса

	http.HandleFunc("/bodyparse", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(data))
		r.Body.Close()
	})

	// basic author
	//через curl:  curl --header "Authorization: Basic YWE6YWFh" "localhost:8080/secure"

	http.HandleFunc("/secure", func(w http.ResponseWriter, r *http.Request) {
		name, pass, ok := r.BasicAuth()
		if !ok {
			_, _ = fmt.Fprintln(w, "aut header not found")
		}
		_, _ = fmt.Fprintln(w, fmt.Sprintf("name: %s, pass: %s", name, pass))

	})

	type User struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		user := User{"Ivan", 22}
		jsonData, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "invalid jason marshal", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(jsonData)

	})

	fmt.Println("server is listening")
	http.ListenAndServe(":8080", nil)
}
