package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var cond sync.Cond

func main() {
	ch := make(chan int, 5)
	// lock := new(sync.RWMutex)
	consumerCond := sync.NewCond(new(sync.Mutex))
	producerCond := sync.NewCond(new(sync.Mutex))

	wg1 := new(sync.WaitGroup)
	for i := 1; i <= 5; i++ {
		wg1.Add(1)
		go producer(wg1, i, ch, producerCond, consumerCond)
	}
	go func() {
		wg1.Wait()
		close(ch)
	}()
	wg2 := new(sync.WaitGroup)
	for i := 1; i <= 5; i++ {
		wg2.Add(1)
		go consumer(wg2, i, ch, consumerCond, producerCond)
	}

	time.Sleep(time.Second)
	consumerCond.Broadcast()
	wg2.Wait()
	fmt.Println("finish")
}

func producer(wg *sync.WaitGroup, idx int, ch chan int, cond, consumerCond *sync.Cond) {
	defer wg.Done()
	product := 1
	for {
		cond.L.Lock()
		if len(ch) == 5 {
			fmt.Printf("生产者 %d wait.\n", idx)
			cond.Wait()
		}
		ch <- product
		fmt.Printf("---第%d生产者，生产：%d\n", idx, product)
		product++
		cond.L.Unlock()
		consumerCond.Signal()
		if product > 100 {
			break
		}
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	}
}

func consumer(wg *sync.WaitGroup, idx int, ch chan int, cond, producerCond *sync.Cond) {
	defer wg.Done()
	for {
		cond.L.Lock()
		if len(ch) == 1 {
			fmt.Printf("消费者 %d wait.\n", idx)
			cond.Wait()
		}
		product, ok := <-ch
		if !ok {
			cond.L.Unlock()
			break
		}
		fmt.Printf("---第%d消费者，消费：%d\n", idx, product)
		cond.L.Unlock()
		producerCond.Signal()
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	}
	fmt.Printf("消费者 %d quit.\n", idx)
}

// func test1() {
// 	ch := make(chan int, 3)
// 	coNum := 2
// 	done := make(chan bool, coNum)
// 	for i := 1; i <= coNum; i++ {
// 		go Consumer(i, ch, done)
// 	}

// 	go Producer(ch)

// 	for i := 1; i <= coNum; i++ {
// 		<-done
// 	}
// }

// func Producer(ch chan<- int) {
// 	for i := 1; i <= 10; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }

// func Consumer(id int, ch <-chan int, done chan bool) {
// 	for val := range ch {
// 		fmt.Printf("consumer id: %d, recv: %d\n", id, val)
// 	}
// 	fmt.Printf("id: %d, closed\n", id)
// 	done <- true
// }
