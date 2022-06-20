package main

//executable package since func main exist

import (
	"./wordz" //   if mod has to be created, can not linc it, coz it is looking in path and root...
	// go mod init 01_gopackages //превращаем проект в модуль
	"fmt"

	"github.com/fatih/color"
	//reusable package
	// go get github.com/fatih/color  //downloading lib in to go work space
	// after downloadng we may import in to our app
)

func main() {

	fmt.Println("hello world")
	color.Red("Hello in red")
	//wordz.Hello
	fmt.Println(wordz.Random())

}
