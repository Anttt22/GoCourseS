package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("working woth mysql")

	db, err := sql.Open("mysql", "root:Super_22@tcp(127.0.0.1:3306)/golangdb")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('kolya', 22)")
	// if err != nil {
	// 	panic(err)
	// }
	// defer insert.Close()

	res, err := db.Query("SELECT `name`, `age` FROM `users`")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("User: %s, age: %d", user.Name, user.Age))

	}

	fmt.Println("connected to mysql")

}
