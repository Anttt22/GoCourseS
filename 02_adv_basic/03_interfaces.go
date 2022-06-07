package main

import "fmt"

//две структуры
type Email struct {
	Address string
}

type Telephone struct {
	Number  int
	Balance int
}

//обьявдяетс интерфейс с сигнатурой методов(имя, приним парам и возвр парам
// все чо соответствует этой сигнатуре будет типом интрерфейса Sender)
type Sender interface {
	Send(msg string) error
}

//обычное обьявление метода и его реализации для структуры
// (но метод той же сигнатуры что и у интерфейса)
func (e *Email) Send(msg string) error { //could be Email
	fmt.Printf("сообщение\"%v\" отправленно по почте на адрес %v \n", msg, e.Address)
	return nil
}

func (t *Telephone) Send(msg string) error {
	fmt.Printf("сообщение\"%v\" отправленно по телефону с номера %v", msg, t.Number)
	return nil
}

//обыкновенны метод принимающий интерфейс
// можно передавать все структуры реализующие интерфейс
// и как следствие являющиеся и типа этого интерфейса тоже
// реализация метода Send будет вызванна той структуры экземпляр которой
// был передан  как параметр
func Notify(s Sender) {
	err := s.Send("notify message")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {

	email := Email{"www.baza"}
	telephone := Telephone{Number: 553377, Balance: 55}
	Notify(&email)
	Notify(&telephone)
}
