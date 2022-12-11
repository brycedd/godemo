package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func concurrent() {
	//base()
	//baseNote()
	//channel()
	//channel2()
	channel3()
	//ticker()
	//runtimeNote()
}

func runtimeNote() {
	runtime.GOMAXPROCS(1)
	go printTenTimes("11")
	printTenTimes("22")
	time.Sleep(time.Second)
}

func printTenTimes(str string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		if "22" == str && i > 2 {
			runtime.Goexit()
		}
		fmt.Println(str)
	}
}

func ticker() {
	newTicker := time.NewTicker(time.Second)
	for num := 0; num < 5; num++ {
		// 读取不接收channel内值
		<-newTicker.C
		fmt.Println(num, "...")
	}
	newTicker.Stop()
	fmt.Println("ticker demo done")
}

func channel3() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 3: channel 3>>>>>>>>>>>>>>>>>>>>>>>>")
	c1 := make(chan int)
	c2 := make(chan int, 10)
	go func(c chan<- int) {
		for i := 0; i < 10; i++ {
			c1 <- i
			time.Sleep(time.Second * 1)
		}
		fmt.Println("执行完成，c1关闭")
		close(c1)
	}(c1)

	go func(c1 <-chan int, c2 chan<- int) {
		times := 0
		for {
			value, isAlive := <-c1
			times++
			if !isAlive {
				fmt.Println("c1 isAlive", isAlive, times)
			}
			if isAlive {
				c2 <- value * value
			} else {
				break
			}
		}
		close(c2)
	}(c1, c2)

	for i := range c2 {
		fmt.Println("for range:", i)
	}

	//for {
	//	num, isAlive := <-c2
	//	if isAlive {
	//		fmt.Println(num)
	//	} else {
	//		break
	//	}
	//}
}

func channel2() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 3: channel 2>>>>>>>>>>>>>>>>>>>>>>>>")
	// make(chan chanType)
	// make(chan chanType, d unit) d 代表缓冲区大小
	// msg := <-c 读， c为channel
	// c <- msg 写channel
	var c chan string
	var w sync.WaitGroup
	c = make(chan string)
	w.Add(1)
	go func() {
		msg := <-c // 一个读channel
		time.Sleep(time.Second * 2)
		fmt.Println("i am reader,", msg)
		w.Done()
	}()
	fmt.Println("开始写入channel")
	c <- "hello" // 写入消息到channel c
	fmt.Println("开始等待goroutine读取结束")
	w.Wait()
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

func baseNote() {
	fmt.Println("begin call goroutine")
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		fmt.Println("goroutine")
		w.Done()
	}()
	w.Wait()
	//time.Sleep(time.Second * 2)
	fmt.Println("end call goroutine")
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
