package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func concurrent2() {
	now := time.Now()
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)
	stop := make(chan struct{})
	w := sync.WaitGroup{}
	senderNum := 10
	jobNum := 100
	w.Add(1)
	createPoll(senderNum, jobChan, resultChan, stop)
	go func(resultChan chan *Result, jobChan chan *Job, w *sync.WaitGroup) {
		defer w.Done()
		stop := 0
		for result := range resultChan {
			fmt.Printf("job id:%v name:%s randum:%v result:%d\n",
				result.job.Id, result.goroutineName, result.job.RandNum, result.sum)
			stop++
			if stop == jobNum {
				return
			}
		}
	}(resultChan, jobChan, &w)
	var id int
	for id < jobNum {
		id++
		rNum := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: rNum,
		}
		jobChan <- job
	}
	w.Wait()
	fmt.Println("end: ", time.Now().Sub(now))
}

func createPoll(num int, jobChan chan *Job, resultChan chan *Result, stop chan struct{}) {
	for i := 0; i < num; i++ {
		namId := i + 1
		go func(jobChan chan *Job, resultChan chan *Result, stop chan struct{}) {
			goroutineName := fmt.Sprintf("%s:%d", "name", namId)
			for job := range jobChan {
				rNum := job.RandNum
				sum := sumNumber(rNum)
				rr := newResultBuilder().
					Value(job, sum, goroutineName).
					Build()
				resultChan <- rr
			}
		}(jobChan, resultChan, stop)
	}
}

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	job           *Job
	sum           int
	goroutineName string
}

func sumNumber(rNum int) int {
	var sum int
	for rNum != 0 {
		tmp := rNum % 10
		sum += tmp
		rNum /= 10
	}
	return sum
}

func newResultBuilder() *ResultBuilder {
	return &ResultBuilder{rr: &Result{}}
}

type ResultBuilder struct {
	rr *Result
}

func (rb *ResultBuilder) Value(job *Job, sum int, name string) *ResultBuilder {
	rb.rr.job = job
	rb.rr.sum = sum
	rb.rr.goroutineName = name
	return rb
}

func (rb *ResultBuilder) Build() *Result {
	return rb.rr
}
