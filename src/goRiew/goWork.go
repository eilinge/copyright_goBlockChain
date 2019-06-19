package main

import (
	"fmt"
	"time"
)

// 编码实现：启动三个协程,第一个协程将0,1,2,3……依次传递给第二个协程,
// 第二个协程将得到的数字平方后传递给第三个协程,第三个协程负责打印得到的数字

// 三个协程(相互通信)
// 一个channal(全局, 不需要作为参数传入到函数)
// chan<- uint8 必须在协程中执行

var chanTrans chan uint8

func transfer() {
	var n1 uint8
	for {
		n := <-chanTrans
		n1 = n * n
		go receive(n1)
	}
}

func receive(n uint8) {
	fmt.Println(n)
}

func main() {
	chanTrans = make(chan uint8)
	defer close(chanTrans)
	var i uint8
	go transfer()
	go func() {
		for i = 0; i < 10; i++ {
			chanTrans <- i
		}
	}()
	time.Sleep(time.Millisecond * 10)
}
