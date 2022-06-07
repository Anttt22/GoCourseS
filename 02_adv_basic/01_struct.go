package main

import (
	"fmt"
)

//STRUCT

type Client struct {
	ID   int64
	Name string
	Age  int
	IMG  *Avatar //поле имг явдяется ссылкой на структуру аватар
} // или, img типа  указатель на AVATAR
type Avatar struct {
	URL  string
	size int64
}

func main() {

	i := new(int64)
	_ = i

	client := Client{}       //не передаем данных все поля будт дефолтного типа
	client.IMG = new(Avatar) // зарезервируем память
	//client.IMG = &Avatar{} //можно и так
	fmt.Printf("%#v\n", client)
	updateAvatar(&client) // передача поссылке
	//updateClient(&client)
	//fmt.Printf("%#v\n", client)
}

func updateAvatar(cl *Client) { //предача по ссылке

	cl.IMG.URL = "new url"
}
func updateClient(cl *Client) {
	cl.Name = "Artem"
}
