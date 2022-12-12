package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func channelDemo() {
	//close1()
	close2()
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

			dataCh <- value
			//select {
			//case <-stopCh:
			//	return
			//case dataCh <- value:
			//}
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
