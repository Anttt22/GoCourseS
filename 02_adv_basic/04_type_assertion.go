package main

import "fmt"

type Email struct {
	Address string
}

type Telephone struct {
	Number  int
	Balance int
}

type Sender interface {
	Send(msg string) error
}

func (e *Email) Send(msg string) error { //could be Email
	fmt.Printf("сообщение\"%v\" отправленно по почте на адрес %v \n", msg, e.Address)
	return nil
}

func (t *Telephone) Send(msg string) error {
	fmt.Printf("сообщение\"%v\" отправленно по телефону с номера %v", msg, t.Number)
	return nil
}

func Notify(s Sender) {
	err := s.Send("notify message")
	if err != nil {
		fmt.Println(err)
		return
	}
	// s. - есть доступ только к методу интерфейса
	// s.Number и другие поля структур недоступны
	// в этом случае можно воспользоваться Утверждением типов
	ttelephone, ok := s.(*Telephone) //если смогли утвердить то сможем вывести
	if ok {
		fmt.Println("\nfrom ok construkcii balance - ", ttelephone.Balance)
	}

	//так же удонее использовать switch case чтобы
	//определить что за тип скрывается под интересом
	// этот вариант использования switch case  называют type switch
	switch s.(type) {
	case *Email:
		fmt.Println("сообщение по почте")
	case *Telephone:
		tel := s.(*Telephone)
		fmt.Println("from switch type: balance =", tel.Balance)
	}

	fmt.Println("sucess")
}

func main() {

	email := Email{"www.baza"}
	telephone := Telephone{Number: 553377, Balance: 55}
	Notify(&email)
	Notify(&telephone)
}
