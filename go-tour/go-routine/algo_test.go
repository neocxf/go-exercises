package main

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{108, 15, 50, 4, 8, 42, 23, 16}

	//bubleSort(arr)

	printArr(mergeSort(arr))
}

func printArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func bubleSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr

}

func mergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	var left, right []int

	left = mergeSort(arr[:mid])
	right = mergeSort(arr[mid:])

	return merge(left, right)
}

func TestMerge(t *testing.T) {
	leftArr := []int{10, 15, 20, 24}
	rightArr := []int{8, 16, 18, 31}

	result := merge(leftArr, rightArr)

	printArr(result)
}

func merge(left []int, right []int) []int {

	result := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}

		if len(right) == 0 {
			return append(result, left...)
		}

		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}

	}

	return result
}

func TestSem(t *testing.T) {
	var sem = make(chan struct{}, 2)

	select {
	case sem <- struct{}{}:
		go func() {
			<-sem
			fmt.Println("worked 1")
		}()
	default:
		fmt.Println("default case 1")
	}

	select {
	case sem <- struct{}{}:
		go func() {
			<-sem
			fmt.Println("worked 2")
		}()
	default:
		fmt.Println("default case 2")
	}

	select {
	case sem <- struct{}{}:
		go func() {
			<-sem
			fmt.Println("worked 3")
		}()
	default: // if case blocked, go to this:
		fmt.Println("default case 3")
	}
}
