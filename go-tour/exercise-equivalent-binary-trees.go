package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func (t *Tree) New(k int) *Tree {

	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}

	return t
}

func insert(tree *Tree, v int) *Tree {
	if tree == nil {
		return &Tree{nil, v, nil}
	}

	if v < tree.Value {
		tree.Left = insert(tree.Left, v)
	} else {
		tree.Right = insert(tree.Right, v)
	}

	return tree
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}

	s := ""

	if t.Left != nil {
		s += t.Left.String() + ", "
	}

	s += fmt.Sprint(t.Value)

	if t.Right != nil {
		s += ", " + t.Right.String()
	}

	return "(" + s + ")"
}

// Walk walks the tree t sending all values from the tree to the channel ch
func Walk(t *Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

func _walk(t *Tree, ch chan int) {
	if t != nil {
		_walk(t.Left, ch)
		ch <- t.Value
		_walk(t.Right, ch)
	}
}

func Same(t1, t2 *Tree) bool {

	return false
}

func main() {

	tree := &Tree{}
	tree.New(1)

	fmt.Println(tree)

	c := make(chan int)

	wg.Add(1)
	go Walk(tree, c)

	for v := range c {
		fmt.Println(v)
	}

}
