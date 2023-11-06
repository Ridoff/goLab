package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int, left, right int) {
	if left < right {
		pivotIndex := partition(arr, left, right)
		quickSort(arr, left, pivotIndex-1)
		quickSort(arr, pivotIndex+1, right)
	}
}

func partition(arr []int, left, right int) int {
	pivotIndex := rand.Intn(right-left+1) + left
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] >= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

func findKthLargest(arr []int, k int) int {
	left, right := 0, len(arr)-1
	for {
		pivotIndex := partition(arr, left, right)
		if pivotIndex == k-1 {
			return arr[pivotIndex]
		} else if pivotIndex < k-1 {
			left = pivotIndex + 1
		} else {
			right = pivotIndex - 1
		}
	}
}

func main() {
	var n, k int
	fmt.Print("Enter the size of the array: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Printf("Enter %d elements separated by spaces: ", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("Enter the value of k: ")
	fmt.Scan(&k)

	quickSort(arr, 0, n-1)
	kthLargest := findKthLargest(arr, k)
	fmt.Printf("The %d-th largest element is: %d\n", k, kthLargest)
}
