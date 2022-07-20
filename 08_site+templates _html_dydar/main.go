package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                 string
	Age                  uint16
	Money                int16
	Avg_grades, Happines float64
	Hobbies              []string
}

func main() {

	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"footbal", "scate", "dance"}}
	// fmt.Fprintf(w, `<h1> Main Text </h1>
	// 	<b>maint text</b>`)
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)

}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts Page")
}

func (u User) getAllinfo() string {
	return fmt.Sprintf("User name is: %s He is %d and He has money %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}
