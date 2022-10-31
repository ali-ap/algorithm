package problems

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}
	foundPos := BinarySearch(nums, 0, len(nums)-1, target)
	if foundPos == -1 {
		return []int{-1, -1}
	}
	first, last, tempFirst, tempLast := foundPos, foundPos, 0, 0

	for first != -1 {
		tempFirst = first
		first = BinarySearch(nums, 0, first-1, target)
	}
	first = tempFirst

	for last != -1 {
		tempLast = last
		last = BinarySearch(nums, last+1, len(nums)-1, target)
	}
	last = tempLast

	return []int{first, last}
}

func SolveSearchRange() {
	fmt.Println("34. Find First and Last Position of Element in Sorted Array")
	fmt.Println("https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/")
	fmt.Println("---------------------------")
}
