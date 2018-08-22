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

func TestP_C(t *testing.T) {
	synChan := make(chan int, 5)
	strChan := make(chan string, 4)

	var wg sync.WaitGroup

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			for elem := range strChan {
				fmt.Println("recieve task: ", elem)
			}

			fmt.Printf("[reciever] %d task done\n", i)
			synChan <- i

			wg.Done()
		}(i)
	}

	go func() {
		for _, elem := range []string{"a", "b", "c", "d", "e", "f"} {
			strChan <- elem
			fmt.Printf("send: %s to [receiver]\n", elem)
		}

		close(strChan)
	}()

	wg.Wait()

	close(synChan)

	for v := range synChan {
		_ = v
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

	var wg sync.WaitGroup
	var wg2 sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			if i, ok := <-queue; ok {
				fmt.Println(i)
			} else {
				break
			}
		}

	}()

	wg2.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg2.Done()
			queue <- i
		}(i)

	}

	go func() {
		wg2.Wait()
		close(queue)
	}()

	wg.Wait()

}

func TestP_C_no_wg(t *testing.T) {
	c := make(chan int)
	signal := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(c chan int, signal chan bool) {

			c <- rand.Intn(10)

			signal <- true

		}(c, signal)
	}

	go func() {
		for {
			if elem, ok := <-c; ok {
				fmt.Println(elem)
			} else {
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		<-signal
	}

}

// without waitGroup, we have to know the exact number of go-routines we spawn
func TestC_P_no_wg(t *testing.T) {
	c := make(chan string) // make the chan unbuffered
	signal := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(c chan string, signal chan bool) {

			defer func() {
				signal <- true
			}()

			for {
				if elem, ok := <-c; ok {
					fmt.Println("doing task: ", elem)
				} else {
					return
				}
			}

		}(c, signal)
	}

	go func() {
		for _, elem := range []string{"a", "b", "c", "d", "e", "f"} {
			fmt.Printf("send: %s to [receiver]\n", elem)
			c <- elem
		}

		close(c)
	}()

	for i := 0; i < 10; i++ { // as we know the number of go-routine we spawn, we just drain the chan
		<-signal
	}

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

func TestWaitSomethingFinish(t *testing.T) {
	done := make(chan struct{})

	go func() {
		time.Sleep(time.Second)
		done <- struct{}{}
	}()

	fmt.Println("go routine executing ...")

	<-done

	fmt.Println("go routine done")
}

func TestStartThingsAtSameTime(t *testing.T) {
	start := make(chan struct{})

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(i int) {
			defer wg.Done()
			<-start

			fmt.Println(i)
		}(i)
	}

	fmt.Println("kick start all the process ...")

	close(start)

	wg.Wait()
}

func TestStopChannelBySignal(t *testing.T) {
	dataChan := make(chan int)
	stopChan := make(chan struct{})

	wg.Add(1)

	go func() {
		for i := 0; i < 10; i++ {
			dataChan <- i
		}

		stopChan <- struct{}{}
	}()

	go func() {
		defer wg.Done()

	loop:
		for {
			select {
			case m := <-dataChan:
				fmt.Println(m)
			case <-stopChan:
				break loop
			}
		}

	}()

	wg.Wait()

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
