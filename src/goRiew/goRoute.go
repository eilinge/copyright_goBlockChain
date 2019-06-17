package main

import (
	"fmt"
	"time"
)

const (
	n = 45
)

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-*/" {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func main03() {
	go spinner(100 * time.Millisecond) // 开启一个协程
	fibN := fib(n)                     // 进程
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
