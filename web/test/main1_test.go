package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestUnbounded(t *testing.T) {
	queue := make(chan string)
	//quitChan := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-queue
	}()
	queue <- "one"
}

func TestBoundedAndReceive(t *testing.T) {
	queue := make(chan string, 2)

	queue <- "one"
	queue <- "two"

	<-queue
	<-queue

	queue <- "three"
}

func TestUnboundedAndReceive(t *testing.T) {
	queue := make(chan string)

	go func() {
		queue <- "one"
	}()

	val := <-queue

	fmt.Println(val)

}

func TestUnboundedAndSend(t *testing.T) {
	queue := make(chan string)

	go func() {
		val := <-queue

		fmt.Println(val)
	}()

	queue <- "one"

}

func TestUnboundedAndSend2(t *testing.T) {
	queue := make(chan string)

	go func() {
		val := <-queue
		fmt.Println(val)
		val = <-queue
		fmt.Println(val)
		val = <-queue
		fmt.Println(val)
	}()

	queue <- "one"
	queue <- "two"
	queue <- "three"
	time.Sleep(time.Second)
}

// we don't use any sync, so the producer may produce data that's not consumed by the consumer
func TestUnboundedAndSend2WithRangeProblematic(t *testing.T) {
	queue := make(chan string)

	go func() {
		for val := range queue {
			fmt.Println(val)
		}
	}()

	queue <- "one"
	queue <- "two"
	queue <- "three"
	queue <- "four" // here we implicitly close the channel

}

func TestUnboundedAndSend2WithRange(t *testing.T) {
	queue := make(chan string)

	go func() {
		defer wg.Done()

		for val := range queue {
			fmt.Println(val)
		}

	}()

	wg.Add(1)

	queue <- "one"
	queue <- "two"
	queue <- "three"
	queue <- "four" // here we implicitly close the channel
	close(queue)

	wg.Wait()
}

func TestUnboundedAndReceiveWithQuit(t *testing.T) {
	queue := make(chan string, 5)
	quit := make(chan bool)

	go func() {

		queue <- "one"
		queue <- "two"
		queue <- "three"
		queue <- "four" // here we implicitly close the channel

		quit <- true
	}()

	<-quit

	close(queue)

	for val := range queue {
		fmt.Println(val)
	}

}

func TestUnboundedAndReceiveWithRangeAndBuffered(t *testing.T) {
	queue := make(chan string)

	go func() {
		queue <- "one"
		queue <- "two"
		queue <- "three"
		queue <- "four" // here we implicitly close the channel

		close(queue)
	}()

	//fmt.Println(<-queue)
	//fmt.Println(<-queue)
	//fmt.Println(<-queue)
	//fmt.Println(<-queue)

	for val := range queue {
		fmt.Println(val)
	}

	//for {
	//	if val, ok := <-queue; ok {
	//		fmt.Println(val)
	//	} else {
	//		break
	//	}
	//}

}

func TestUnboundedAndSend2WithRangeAnd2GoRoutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		queue := make(chan string)

		count := 0

		wg.Add(1)

		go func() {
			defer wg.Done()

			for val := range queue {
				fmt.Println(val)
				count++
			}

		}()

		go func() {
			queue <- "one"
			queue <- "two"
			queue <- "three"
			queue <- "four" // here we implicitly close the channel
			close(queue)
		}()

		wg.Wait()

		assert.Equal(t, count, 4)
	}
}

func TestUnboundedAndSend2WithRangeAnd2GoRoutineAndProducerSlow(t *testing.T) {
	for i := 0; i < 5; i++ {
		queue := make(chan int)

		count := 0

		wg.Add(1)

		go func() {
			defer wg.Done()

			for val := range queue {
				fmt.Println(val)
				count++
			}

		}()

		go func() {
			for i := 0; i < 4; i++ {
				queue <- i
				time.Sleep(time.Second)
			}
			close(queue)
		}()

		wg.Wait()

		assert.Equal(t, count, 4)
	}
}

func TestUnboundedAndSend2WithForAnd2GoRoutineAndProducerSlow(t *testing.T) {
	for i := 0; i < 5; i++ {
		queue := make(chan int)

		count := 0

		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				if i, ok := <-queue; ok { // we can id the status of the channel by receiving it with .ok
					count++
					fmt.Println(i)
				} else {
					break
				}
			}

		}()

		go func() {
			for i := 0; i < 4; i++ {
				queue <- i
				//time.Sleep(time.Second)
			}
			close(queue)
		}()

		wg.Wait()

		assert.Equal(t, count, 4)
	}
}

func TestUnboundedAndSend2WithSelectAnd2GoRoutineAndProducerSlow(t *testing.T) {
	queue := make(chan int)
	quitCh := make(chan struct{})

	go func() {
		for {
			select {
			case elem := <-queue:
				fmt.Println(elem)
			case <-quitCh:
				break
			}
		}

	}()

	go func() {
		for i := 0; i < 4; i++ {
			queue <- i
			//time.Sleep(time.Second)
		}
		//close(queue)
	}()

	time.Sleep(time.Second)

	quitCh <- struct{}{}
}

func TestProducerAndConsumer(t *testing.T) {
	for j := 0; j < 5; j++ {
		queue := make(chan int)

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				queue <- i
			}(i)
		}

		counter := 0

		go func() {
			for {
				if i, ok := <-queue; ok {
					counter++
					fmt.Println(i)
				} else {
					break
				}
			}

		}()

		wg.Wait()

		assert.Equal(t, counter, 10)
	}

}

func TestProducerAndConsumerWithSleep(t *testing.T) {

	for i := 0; i < 1000; i++ {
		queue := make(chan int)

		counter := 0

		for i := 0; i < 100; i++ {
			go func(i int) {
				queue <- i
				time.Sleep(time.Millisecond * 100)
			}(i)
		}

		go func() {
			for i := range queue {
				counter++
				fmt.Println(i)
			}

		}()

		time.Sleep(time.Millisecond * 10)
		assert.Equal(t, counter, 100)
	}

}

func TestProducerAndConsumerWithWg(t *testing.T) {

	for i := 0; i < 1000; i++ {
		var wg sync.WaitGroup
		queue := make(chan int)

		counter := 0

		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				fmt.Println("sending")
				queue <- i
				fmt.Println("sent")
			}(i)
		}

		go func() {
			for i := range queue {
				counter++
				fmt.Println("consume: ", i)
			}

		}()

		wg.Wait()

		assert.Equal(t, counter, 100)
	}

}

func TestProduceAndConsumer(t *testing.T) {
	for i := 0; i < 1; i++ {
		queue := make(chan int)
		done := make(chan bool)

		go func() {
			for i := 0; i < 10; i++ {
				fmt.Println("sending")
				queue <- i
				fmt.Println("sent")
			}

			fmt.Println("before closing channel")
			close(queue)
			fmt.Println("before passing true to done")
			done <- true
		}()

		go func() {
			for i := range queue {
				fmt.Println("consume: ", i)
				time.Sleep(time.Millisecond * 100)
			}
		}()

		<-done

		fmt.Println("After calling DONE")
	}
}

func TestProducerAndConsumerUsingSignals(t *testing.T) {
	c := make(chan int)
	signal := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(c chan int, signal chan bool) {

			c <- rand.Int()

			signal <- true

		}(c, signal)
	}

	//go func() {
	//	for {
	//		select {
	//		case num := <-c:
	//			fmt.Println("consumed ", num)
	//		case <-quit:
	//			break
	//		}
	//	}
	//}()

	go func() {
		for i := range c {
			fmt.Println("consumed ", i)
		}
	}()

	for i := 0; i < 10; i++ { // drain the pool to confirm that all work is done
		<-signal
	}

	fmt.Println("all done.")

}

func TestProducerAndConsumerUsingQuitWorksInBuffered(t *testing.T) {
	c := make(chan int, 10)
	signal := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(c chan int, signal chan bool) {

			c <- rand.Int()

			signal <- true

		}(c, signal)
	}

	for i := 0; i < 10; i++ { // drain the pool to confirm that all work is done
		<-signal
	}

	fmt.Println("called close")

	close(c)

	for i := range c {
		fmt.Println("consumed ", i)
	}

	fmt.Println("all done.")

}

func TestProducerAndConsumerUsingQuitWorksInUnBuffered(t *testing.T) {
	c := make(chan int)
	signal := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(c chan int, signal chan bool) {

			c <- rand.Int()

			signal <- true

		}(c, signal)
	}

	for i := range c {
		fmt.Println("consumed ", i)
	}

	fmt.Println("all done.")

}

func TestOneProducerAndConsumerUsingQuitWorksInUnBuffered(t *testing.T) {
	c := make(chan int)
	signal := make(chan bool)

	go func(c chan int, signal chan bool) {
		for i := 0; i < 10; i++ {
			c <- rand.Int()
		}

		close(c)

	}(c, signal)

	for i := range c {
		fmt.Println("consumed ", i)
	}

	fmt.Println("all done.")

}

func TestProducerAndConsumerUsingWaitGroup(t *testing.T) {
	var producer_wg sync.WaitGroup
	var consumer_wg sync.WaitGroup

	c := make(chan int)

	for i := 0; i < 10; i++ {
		producer_wg.Add(1)

		go func(c chan int) {
			defer producer_wg.Done()
			c <- rand.Int()

		}(c)
	}

	consumer_wg.Add(1)
	go func() {
		defer consumer_wg.Done()
		for i := range c {
			fmt.Println("consumed ", i)
		}
	}()

	producer_wg.Wait()

	close(c)

	consumer_wg.Wait()

	fmt.Println("all done.")

}

func TestSimpleConsume(t *testing.T) {
	c := make(chan bool)

	go func() {
		<-c
		<-c
	}()

	c <- true
	c <- true
}
