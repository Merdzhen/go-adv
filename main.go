package main

import (
	"fmt"
	"net/http"
	"time"
)

// написать функцию которая делает 10 конкурентных запросов на get по google.com
// вывести в консоль 10 statuscode

func main() {
	t := time.Now()
	for i := 0; i < 10; i++ {
		go getHttpStatus()
	}
	time.Sleep(time.Millisecond * 1500)
	fmt.Println(time.Since(t))
}

func getHttpStatus() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Printf("Ошибка %s", err.Error())
		return
	}
	fmt.Printf("Код %d \n", resp.StatusCode)
}
