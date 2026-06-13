package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// написать функцию которая делает 10 конкурентных запросов на get по google.com
// вывести в консоль 10 statuscode

func main() {
	t := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go getHttpStatus(&wg)
	}
	wg.Wait()
	fmt.Println(time.Since(t))
}

func getHttpStatus(wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Printf("Ошибка %s", err.Error())
		return
	}
	fmt.Printf("Код %d \n", resp.StatusCode)
}
