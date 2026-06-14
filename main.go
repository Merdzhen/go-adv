package main

import (
	"fmt"
	"net/http"
)


func main() {
	code := make(chan int)
	go getHttpStatus(code)
	<- code
}

func getHttpStatus(codeCh chan int) {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Printf("Ошибка %s", err.Error())
		return
	}
	fmt.Printf("Код %d \n", resp.StatusCode)
	codeCh <- resp.StatusCode
}
