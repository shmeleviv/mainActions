package main

import (
	"fmt"
)

func InsertionSort(A []int) {
	key := 0
	i := 0
	for j := 1; j < len(A); j++ {
		key = A[j]
		i = j - 1
		fmt.Printf("key = %d\n", key)
		for i >= 0 && A[i] > key {
			A[i+1], A[i] = A[i], A[i+1]
			fmt.Println(A)
			i -= 1

		}

	}
}

func InsertionSortReverse(A []int) {
	key := 0
	i := 0
	for j := 1; j < len(A); j++ {
		key = A[j]
		i = j - 1
		fmt.Printf("key = %d\n", key)
		for i >= 0 && A[i] < key {
			A[i+1], A[i] = A[i], A[i+1]
			fmt.Println(A)
			i -= 1

		}

	}
}

func main() {
	list := []int{31, 41, 59, 26, 41, 58}
	fmt.Println("Start array is", list, "\n")
	InsertionSort(list)
	fmt.Println("Sorted array is", list, "\n")

	list1 := []int{31, 41, 59, 26, 41, 58}
	InsertionSortReverse(list1)
	fmt.Println("Reverse Sorted array is", list1, "\n")

}
