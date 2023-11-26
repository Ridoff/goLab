package main

import (
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