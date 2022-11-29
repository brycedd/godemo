package main

import (
	"fmt"
	"sync"
	"time"
)

func concurrent() {
	//base()
	//channel()
	channel2()

}

func channel2() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 3: channel 2>>>>>>>>>>>>>>>>>>>>>>>>")
	// make(chan chanType)
	// make(chan chanType, d unit) d 代表缓冲区大小
	// msg := <-c 读， c为channel
	// c <- msg 写channel
	var c chan string
	var w2 sync.WaitGroup
	c = make(chan string)
	w2.Add(1)
	go func() {
		msg := <-c // 一个读channel
		time.Sleep(time.Second * 2)
		fmt.Println("i am reader,", msg)
		w2.Done()
	}()
	fmt.Println("开始写入channel")
	c <- "hello" // 写入消息到channel c
	fmt.Println("开始等待goroutine读取结束")
	w2.Wait()
}

func channel() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 2: channel>>>>>>>>>>>>>>>>>>>>>>>>")
	//w := sync.WaitGroup{}
	var w sync.WaitGroup
	for i := 0; i < 5; i++ {
		// 增加一个计数
		w.Add(1)
		fmt.Printf("waitGroup add %d finish  ", i)
		go func(num int) {
			time.Sleep(time.Second * time.Duration(num))
			fmt.Printf("\ni am %d Goroutine", num)
			// 减少一个计数
			w.Done()
		}(i)
	}
	// 阻塞主线程，知道waitGroup计数变为为0
	w.Wait()
}

func base() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 0: goroutine>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("begin call goroutine")
	go func() {
		fmt.Println("i am a goroutine")
	}()
	fmt.Println(">>>>>>>>>>>>>>>>>>part 1: fib()>>>>>>>>>>>>>>>>>>>>>>>>")
	go spinner(time.Millisecond * 1000)
	// 下一行代码执行完，main线程将关闭，上一行的goroutine也会结束
	fmt.Printf("\n%d\n", fib(45))
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-2) + fib(x-1)
}
