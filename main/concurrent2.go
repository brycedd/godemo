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
	w := sync.WaitGroup{}
	senderNum := 10
	jobNum := 100
	w.Add(jobNum)
	createPoll(senderNum, jobChan, resultChan)
	go func(resultChan chan *Result, jobChan chan *Job) {
		for result := range resultChan {
			fmt.Printf("job id:%v name:%s randum:%v result:%d\n",
				result.job.Id, result.goroutineName, result.job.RandNum, result.sum)
			w.Done()
		}
	}(resultChan, jobChan)
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
	fmt.Println("waiting...")
	w.Wait()
	fmt.Println("end: ", time.Now().Sub(now))
}

func createPoll(num int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < num; i++ {
		namId := i + 1
		go func(jobChan chan *Job, resultChan chan *Result) {
			goroutineName := fmt.Sprintf("%s:%d", "name", namId)
			for job := range jobChan {
				rNum := job.RandNum
				sum := sumNumber(rNum)
				rr := newResultBuilder().
					Value(job, sum, goroutineName).
					Build()
				resultChan <- rr
			}
		}(jobChan, resultChan)
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
