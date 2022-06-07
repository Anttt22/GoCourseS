package main

import (
	"fmt"
)

// Metods
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

// одьявление метода для структуры Client
func (c Client) HasAvatar() bool {
	if c.IMG != nil && c.IMG.URL != "" {
		return true
	} else {
		return false
	}
}

func (c Client) UpdateAvatar(u string) {
	c.IMG.URL = u
}

func main() {

	client := Client{
		ID:   33,
		Name: "Artiom",
		Age:  22,
	}
	fmt.Println(client.HasAvatar())

	client2 := Client{
		ID:   011,
		Name: "BondDick",
		Age:  79,
		IMG: &Avatar{
			URL:  "www.dog.com",
			size: 256,
		}, //с помощью обьявления пустой структуры и порлучения ссылки на нее
		//IMG:  new(Avatar),  // или так

	}
	fmt.Println(client2.HasAvatar())
	client2.UpdateAvatar("www.Teper_KOlya")
	fmt.Println(client2.IMG.URL)
}
