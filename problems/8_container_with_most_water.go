package problems

import (
	"fmt"
)

func maxArea(height []int) int {
	p1, p2, maxArea := 0, len(height)-1, 0

	for p1 < p2 {
		h := Min(height[p1], height[p2])
		w := p2 - p1
		area := w * h
		maxArea = Max(area, maxArea)
		if height[p1] < height[p2] {
			p1++
		} else {
			p2--
		}
	}

	return maxArea
}

func SolveMaxWater() {
	fmt.Println("max area")
	fmt.Println("https://leetcode.com/problems/container-with-most-water/")
	fmt.Println("---------------------------")

	nums1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf(" numbers:  %v >>>  result:  %v \n", nums1, maxArea(nums1))
	fmt.Println("---------------------------")

	nums2 := []int{1, 1}
	fmt.Printf("numbers:  %v >>>  result:  %v \n", nums2, maxArea(nums2))
	fmt.Println("---------------------------")

}
