package main

import (
	"math/rand"
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

func generateRandomArray(size, maxValue int) []int {
    arr := make([]int, size)
    for i := 0; i < size; i++ {
        arr[i] = rand.Intn(maxValue)
    }
    return arr
}

func main() {
    var k, size int
    _, err := fmt.Scan(k)
    if err == nil {
            fmt.Println("Error:", err)
    }
    _, erm := fmt.Scan(size)
    if erm == nil {
            fmt.Println("Error:", erm)
    }
    rand.Seed(time.Now().UnixNano())
    maxValue := 100
    arr := generateRandomArray(size, maxValue)
    fmt.Println("Generated array:", arr)

    result := findKthLargest(arr, k)

    fmt.Printf("The %dth largest element is: %d\n", k, result)
}