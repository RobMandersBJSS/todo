package main

import (
	"fmt"
)

func main() {
	channelOdd := make(chan int)
	channelEven := make(chan int)
	num := 0

	go updateOdd(channelOdd)
	go updateEven(channelEven)

	for i := 0; i < 10; i++ {
		select {
		case odd := <-channelOdd:
			num = odd
			fmt.Println(num)
		case even := <-channelEven:
			num = even
			fmt.Println(num)
		}
	}
}

func updateOdd(channel chan int) {
	for i := 1; i < 10; i += 2 {
		channel <- i
	}
}

func updateEven(channel chan int) {
	for i := 2; i <= 10; i += 2 {
		channel <- i
	}
}
