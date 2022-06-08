package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	////another example
	// var answer1, answer2, answer3 string

	// fmt.Print("Name: ")
	// _, err = fmt.Scan(&answer1)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Print("Fav food: ")
	// _, err = fmt.Scan(&answer2)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Print("Fav sport: ")
	// _, err = fmt.Scan(&answer3)
	// if err != nil {
	// 	panic(err)
	// }

	////and another example
	f, err1 := os.Create("names.txt")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer f.Close()
	r := strings.NewReader("bonobono")
	io.Copy(f, r)

}
