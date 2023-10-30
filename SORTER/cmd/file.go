package main

import (
    "fmt"
)
func findKthLargest(nums []int, k int) int {
    quickSort(nums, 0, len(nums)-1)
    return nums[len(nums)-k]
}

func quickSort(nums []int, left, right int) {
    if left < right {
        pivotIndex := partition(nums, left, right)
        quickSort(nums, left, pivotIndex-1)
        quickSort(nums, pivotIndex+1, right)
    }
}

func partition(nums []int, left, right int) int {
    pivot := nums[right]
    i := left - 1

    for j := left; j < right; j++ {
        if nums[j] >= pivot {
            i++
            nums[i], nums[j] = nums[j], nums[i]
        }
    }

    nums[i+1], nums[right] = nums[right], nums[i+1]
    return i + 1
}

func main() {
    arr := []int{3, 2, 1, 5, 6, 4}
    k := 2;
    result := findKthLargest(arr, k)
    fmt.Printf("The %dth largest element is: %d\n", k, result)
}