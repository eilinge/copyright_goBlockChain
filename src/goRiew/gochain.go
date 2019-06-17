package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func sendNum(c chan int) {
	fmt.Println("begin call send NuM")
	c <- 6
	fmt.Println("send Num is end")
}

func main04() {
	var c = make(chan int)
	defer close(c)
	go sendNum(c)
	time.Sleep(time.Second)
	// num := <-c
	// begin call send NuM
	// the int in chan:  6
	// send Num is end
	fmt.Println("the int in chan: ", <-c)
	// begin call send NuM
	// the int in chan:  6
	runtime.Gosched() // 保证另外的协程彻底结束
}

func lunch() {
	fmt.Println("发射")
}

func main() {
	abort := make(chan string)
	go func() {
		buf := make([]byte, 1)

		os.Stdin.Read(buf)
		abort <- "cancel"
	}()

	tk := time.NewTicker(time.Second * 1)
	defer tk.Stop()
	// for {
	for i := 5; i > 0; i-- {
		fmt.Println(i)
		select {
		case <-tk.C:
		case num := <-abort:
			fmt.Println(num)
			return
		}
	}
	// }
	lunch()

}
