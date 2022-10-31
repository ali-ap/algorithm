package problems

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, indexToFind int) int {
	index := partition(nums, left, right)
	if index == indexToFind {
		return nums[index]
	} else if indexToFind < index {
		return quickSelect(nums, left, index-1, indexToFind)
	} else {
		return quickSelect(nums, index+1, right, indexToFind)
	}
}

func partition(nums []int, left, right int) int {
	i := left
	for j := left; j <= right; j++ {
		if nums[j] <= nums[right] {
			temp := nums[j]
			nums[j] = nums[i]
			nums[i] = temp
			i++
		}
	}
	return i - 1
}

func SolveFindKthLargest() {
	fmt.Println("215. Kth Largest Element in an Array")
	fmt.Println("https://leetcode.com/problems/kth-largest-element-in-an-array/")
	fmt.Println("---------------------------")
}
