package main

import "fmt"

//EMPTY INTERFACE

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

type Caller interface {
	Call(num int) error
}

type Iphone interface {
	Sender
	Caller
	Mymethod()
	MyMethod2()
}

func (e *Email) Send(msg string) error { //could be Email
	fmt.Printf("сообщение\"%v\" отправленно по почте на адрес %v \n", msg, e.Address)
	return nil
}

func (t *Telephone) Send(msg string) error {
	fmt.Printf("сообщение\"%v\" отправленно по телефону с номера %v", msg, t.Number)
	return nil
}

//изменим тип принимаемого интерфейса напустой интерфейс
//func Notify(s Sender) {    //было
func Notify(i interface{}) {

	switch i.(type) {
	case int: // если получилось привести к инту то выводим сообщение
		fmt.Println("число не поддерживается")
	}
	//у нас еть интерфейс сендер. поэтому можем утвердить тип к этому интерфейсу
	s, ok := i.(Sender)
	if !ok {
		fmt.Println("ошибка утверждения интрефейса")
		return
	}
	err := s.Send("Сообщение пустого интерфейса")
	if err != nil {
		fmt.Println("Ошибка из пустого интерфейса")
		return
	}

	fmt.Println("sucess")
}

func main() {

	//email := Email{"www.baza"}
	//telephone := Telephone{Number: 553377, Balance: 55}
	//Notify(&email)
	//Notify(&telephone)
	//так как метод принимает пустой интерфейс то
	//в него можно отправить все что угодно
	Notify(2)        ///число не поддержвается
	Notify("stroka") ///ошибка утверждения интрефейса
	//даже массив
	arr := [5]int64{1, 5, 2, 3, 4} ///ошибка утверждения интрефейса
	Notify(arr)

}
