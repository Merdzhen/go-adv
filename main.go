package main

import (
	"fmt"
	"net/http"
)


func main() {
	code := make(chan int)
	for i := 0; i < 10; i++ {
		go getHttpStatus(code)
	}
	for res := range code {
		fmt.Printf("Код %d \n", res)
	}
}

func getHttpStatus(codeCh chan int) {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Printf("Ошибка %s", err.Error())
		return
	}
	codeCh <- resp.StatusCode
}
