package main

import (
	"fmt"
	"sync"
)


func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	numGoroutines := 3
	size := len(arr) / numGoroutines
	sumCh :=  make(chan int)
	var wg sync.WaitGroup 
	result := 0

	for i := range numGoroutines {
		start := i * size
		end := start + size
		if i == numGoroutines - 1 {
			end = len(arr)
		}
		wg.Add(1)
		go func(subArr []int) {
			defer wg.Done()
			getSum(subArr, sumCh)
		}(arr[start:end])
	}

	go func() {
		wg.Wait()
		close(sumCh)
	}() 

	for sum := range sumCh {
		result += sum
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
