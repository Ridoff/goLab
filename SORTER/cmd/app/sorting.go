package main

import (
	"math/rand"
)

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	pivot := nums[len(nums)/2]
	left, right := make([]int, 0), make([]int, 0)
	equal := []int{}

	for _, num := range nums {
		switch {
		case num < pivot:
			left = append(left, num)
		case num == pivot:
			equal = append(equal, num)
		case num > pivot:
			right = append(right, num)
		}
	}

	quickSort(left)
	quickSort(right)

	copy(nums, append(append(left, equal...), right...))
}

func generateRandomArray(length int) []int {
	nums := make([]int, length)
	for i := 0; i < length; i++ {
		nums[i] = rand.Intn(100) // Здесь 100 - максимальное случайное число
	}
	return nums
}