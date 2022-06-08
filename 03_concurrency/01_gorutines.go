package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go WorkAnsPrint(i) //запускается горутина и цикл
		// идет на след итерацию не ждя концы выполнения функции
	}

	time.Sleep(100 * time.Millisecond)
}

func WorkAnsPrint(num int) {
	fmt.Println("Job No", num, "started")
	var result int
	for i := 0; i < 1000; i++ {
		result = i * num
	}
	runtime.Gosched() // метод явно говорит плагтровщику
	//что пора переключаться на другую горутину
	fmt.Println("Job No", num, "finished, result = ", result)
}
