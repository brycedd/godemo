package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func channelDemo() {
	//close1()
	//close2()
	blockTest()
}

func blockTest() {
	var w sync.WaitGroup
	w.Add(1)
	c1 := make(chan int, 100)
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
		//close(c1)
		//fmt.Println("c1 closed...")
		fmt.Println("c1 发送结束，但并未手动关闭")
		time.Sleep(time.Second * 3)
		fmt.Println("c1 goroutine 结束休眠")
	}()
	go func() {
		defer w.Done() // c1 未被关闭，不会跳出range，导致这一行永远阻塞，直到main goroutine结束
		for v := range c1 {
			fmt.Println("读取到：", v)
		}
		fmt.Println("退出了for range...")
	}()
	w.Wait()
	fmt.Println("main done")
}

// SafeSend 安全的获取当前channel是否已关闭
func SafeSend(ch chan int, value int) (closed bool) {
	defer func() {
		if recover() != nil {
			// 存在异常
			closed = true
		}
	}()
	ch <- value
	return false
}

func SafClose(ch chan int) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()
	close(ch)
	return true
}

// 优雅的关闭channel方式⬇️
func close1() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const MaxRandomNumber = 10000
	const NumReceivers = 100

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	dataCh := make(chan int, 100)

	// sender
	go func() {
		for {
			// rand.Intn() 生成随机数，相同种子生成器生成结果相同
			if value := rand.Intn(MaxRandomNumber); value == 0 {
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	//receiver
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()
			// receive values until dataCh is closed and
			// the value buffer queue of dataCh is empty
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}
	wgReceivers.Wait()
}

// 一个receiver， 多个sender，利用额外的signal channel实现停止发送
func close2() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const MaxRandomNumber = 1000
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			value := rand.Intn(MaxRandomNumber)

			select {
			case <-stopCh:
				return
			case dataCh <- value:
			}
		}()
	}

	// receiver
	go func() {
		defer wgReceivers.Done()
		for value := range dataCh {
			if value == MaxRandomNumber-1 {
				close(stopCh)
				return
			}
			log.Println(value)
		}
	}()
	wgReceivers.Wait()
}
