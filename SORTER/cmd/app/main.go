package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	if k > len(nums) || k < 1 {
		panic("k is out of range")
	}

	quickSort(nums)
	return nums[len(nums)-k]
}

func findKthSmallest(nums []int, k int) int {
	if k > len(nums) || k < 1 {
		panic("k is out of range")
	}

	quickSort(nums)
	return nums[k-1]
}

func main() {
    var arrSize, k int
    _, err := fmt.Scan(&arrSize)
    _, erm := fmt.Scan(&k)
    if err == nil || erm == nil{
        fmt.Println("Error")
    }
	nums := generateRandomArray(int(arrSize))


	kthLargest := findKthLargest(nums, k)
	kthSmallest := findKthSmallest(nums, k)

	fmt.Printf("K-th Largest: %d\n", kthLargest)
	fmt.Printf("K-th Smallest: %d\n", kthSmallest)
}
