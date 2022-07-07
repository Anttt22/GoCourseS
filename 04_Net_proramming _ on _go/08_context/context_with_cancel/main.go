package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx) //создаем контекс с возможностью отмены
	go func() {
		err := cancelRequest(ctx) //запускаем, ждем 100мсек, возвр ошибку
		if err != nil {           // как только вернулась ошибка вызываем cancel нашего контекста
			cancel()
		}
	}()

	doRequest(ctx, "https://ya.ru")
}

func cancelRequest(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return fmt.Errorf("fail request")
}
func doRequest(ctx context.Context, reqStr string) {
	req, _ := http.NewRequest(http.MethodGet, reqStr, nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//panic(err)
		fmt.Println(err) //можно просто вывести ошиибку. в этом случае
		// распечатается ошибка. потом попадаем в кейс done так как канал закрыт
	}

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("response completed status code= %d", resp.StatusCode)
	case <-ctx.Done():
		//канал закрывается когда работы выполненная для контекста
		// должна быть отменена. Тоесть если работа отменена то канал
		// закрывается
		fmt.Println("request to long")
	}
}
