package main

import (
	"fmt"
)


func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	numGoroutines := 3
	size := len(arr) / numGoroutines
	sumCh :=  make(chan int)
	result := 0

	for i := range numGoroutines {
		start := i * size
		end := start + size
		if i == numGoroutines - 1 {
			end = len(arr)
		}
		go getSum(arr[start:end], sumCh)
	}

	for range numGoroutines {
		result += <-sumCh
	}

	fmt.Println(result)

}

func getSum(arr []int, sumCh chan int) {
	var sum int
	for _, num := range arr {
		sum += num
	}
	sumCh <- sum
}
