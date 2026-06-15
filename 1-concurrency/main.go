package main

import (
	"fmt"
	"math/rand/v2"
)


func main() {
	numCh := make(chan int, 10)
	squareNumCh := make(chan int, 10)
	go createNums(numCh) 
	go squareNums(numCh, squareNumCh)
	// for i := 0; i < 10; i++ {
	// 	squareNum := <- squareNumCh
	// 	fmt.Println(squareNum)
	// }
	for squareNum := range squareNumCh {
		fmt.Println(squareNum)
	}
}

func createNums(numCh chan int) {
	for i := 0; i < 10; i++ {
		numCh <- rand.IntN(101)
	}
	close(numCh)
}

func squareNums(numCh chan int, squareNumCh chan int) {
	// go createNums(numCh) 
	// for i := 0; i < 10; i++ {
	// 	num := <- numCh
	// 	squareNumCh <- num * num
	// }
	for num := range numCh { // Цикл сам остановится, когда numCh закроется
		squareNumCh <- num * num
	}
	close(squareNumCh)
}
