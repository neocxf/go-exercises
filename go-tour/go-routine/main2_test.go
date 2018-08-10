package main

import (
	"fmt"
	"sync"
	"testing"
)

// when you try to read from a closed channel, the channel always return a value, which is the zero value for the channel
func Test2_1(t *testing.T) {

}

// the fist output parameter is the in channel that is write-only( and closable)
// the second output parameter is the read-only out channel
func MakeInfinite() (chan<- interface{}, <-chan interface{}) {

	in := make(chan interface{})
	out := make(chan interface{})

	go func() {
		var inQueue []interface{}
		outChan := func() chan interface{} {
			if len(inQueue) == 0 {
				return nil
			}
			return out
		}
		curVal := func() interface{} {
			if len(inQueue) == 0 {
				return nil
			}

			return inQueue[0]
		}

		for len(inQueue) > 0 || in != nil {
			select {
			case v, ok := <-in:
				if !ok {
					in = nil // by setting in to nil, the select statement will never try to read from in again
				} else {
					inQueue = append(inQueue, v)
				}

				// do something that makes a value go on to the out channel
				// writing to nil channel blocks forever
			case outChan() <- curVal():
				inQueue = inQueue[1:]
			}
		}

		close(out)
	}()

	return in, out
}

func TestMakeInfiniteNoPause(t *testing.T) {
	in, out := MakeInfinite()

	lastVal := -1
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for v := range out {
			vi := v.(int)
			fmt.Println(vi)
			if lastVal+1 != vi {
				t.Errorf("Unexpected value; expected %d, got %d", lastVal+1, vi)
			}

			lastVal = vi
		}

	}()

	for i := 0; i < 1000; i++ {
		fmt.Println("writing", i)
		in <- i
	}

	close(in)

	fmt.Println("finished writing")
	wg.Wait()

	if lastVal != 999 {
		t.Errorf("didn't get all values, last one received was %d", lastVal)
	}
}
