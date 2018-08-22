package wokerpool

import (
	"fmt"
	"sort"
	"testing"
)

func TestGen(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5}

	intChan := gen(arr...)

	for in := range intChan {
		fmt.Println(in)
	}

	c, ok := <-intChan

	if ok {
		fmt.Println(c)
	} else {
		fmt.Println("error")
	}

}

func TestSq(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	intChan := gen(arr...)

	sqChan := sq2(intChan)

	for sq := range sqChan {
		fmt.Println(sq)
	}

	for n := range sq2(sq2(gen(3, 2, 1))) {
		fmt.Println(n)
	}

}

func TestOut1(t *testing.T) {
	in := gen(1, 2, 3)

	sq1 := sq2(in)
	sq2 := sq2(in)

	out := merge2(sq1, sq2)

	fmt.Println(<-out)
	fmt.Println(<-out)

}

func TestOut(t *testing.T) {
	in := gen(1, 2, 3)

	done := make(chan struct{})
	defer close(done)

	sq1 := sq(done, in)
	sq2 := sq(done, in)

	out := merge(done, sq1, sq2)

	fmt.Println(<-out)
	fmt.Println(<-out)
	//fmt.Println(<-out)

	done <- struct{}{}

}

func TestMD5All(t *testing.T) {
	//m, err := MD5All(os.Args[1])
	m, err := MD5All("..")

	if err != nil {
		fmt.Println(err)
		return
	}

	var paths []string

	for path := range m {
		paths = append(paths, path)
	}

	sort.Strings(paths)

	for _, p := range paths {
		fmt.Printf("%x %s\n", m[p], p)
	}
}
