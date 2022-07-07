package main

import (
	"fmt"
	"net/http"
	"time"
)

type ResponseHandler struct {
	message string
}

func (rh ResponseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rh.message)
}

type ResponseHandler2 struct {
	message string
}

func (rh2 ResponseHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, rh2.message)
}

type ResponseHandler3 struct {
	message string
}

type SourceHandler struct {
	message string
}

func (rh3 ResponseHandler3) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	fmt.Fprintf(w, "Request timeout")
}

func (sh SourceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	http.FileServer(http.Dir("./"))
	fmt.Fprintf(w, "filesss")
}

func main() {
	// another way
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "hello")
	// })

	//resp := ResponseHandler{"halloo"}                 //
	//http.Handle("/halloo", resp)                      //on adress/halloo show resp(implements interface)
	//http.Handle("/", http.FileServer(http.Dir("./"))) // returns files
	//http.Handle("/redirect", http.RedirectHandler("https://google.com", 301))

	//Задание

	rhhello := ResponseHandler2{"hello human"}
	rhtimeout := ResponseHandler3{"request timeout"}
	http.Handle("/hello", rhhello)
	http.Handle("/longPing", rhtimeout)

	//file srver for root - easy
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	// not for root:
	//works with strip prefixes/!!!!

	http.Handle("/source/", http.StripPrefix("/source/", http.FileServer(http.Dir("./"))))

	http.ListenAndServe(":8080", nil)

}

// Задание:
// В этом задании вам надо будет поднять http-сервер и зарегистрировать несколько
// обработчиков:
// ● /hello обработчик должен просто поздороваться с нами, отправив сообщение hello
// human;
// ● /source обработчик должен отдавать в качестве ответа файлы из текущей
// директории;
// ● /longPing обработчик должен имитировать долгую работу:
// ○ Обработчик должен быть зарегистрирован через http.Handle.
// ○ Обработчик должен иметь тайм-аут на свое выполнение в 1 секунду, то есть
// через секунду ожидания в ответ должно прийти Request timeout.
// Порядок действий:
// 1. Форкните репозиторий module09 с кодом данного задания - в группу с вашими
// репозиториями - golang_users_repos/<your_gitlab_id>.
// 2. В вашем проекте module09 сделайте новую ветку 01_task.
// 3. В пакете module09/cmd/app зарегистрируйте все нужные обработчики и создайте
// сервер.
// 4. Проверьте работоспособность.
// 5. В ответе пришлите ссылку на merge request в ветку master своего проекта ветки
// 01_task
