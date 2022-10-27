package problems

import "fmt"

func twoSum(nums []int, target int) []int {

	remain := make(map[int]int)

	for k, v := range nums {
		subtract := target - v
		if r, ok := remain[v]; ok {
			return []int{r, k}
		} else {
			remain[subtract] = k
		}
	}
	return []int{}
}

func SolveTwoSum() {
	fmt.Println("two sum")
	fmt.Println("https://leetcode.com/problems/two-sum/")
	fmt.Println("---------------------------")

	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Printf("target:  %v numbers:  %v >>>  result:  %v \n", target1, nums1, twoSum(nums1, target1))
	fmt.Println("---------------------------")

	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Printf("target:  %v numbers:  %v >>>  result:  %v \n", target2, nums2, twoSum(nums2, target2))
	fmt.Println("---------------------------")

	nums3 := []int{3, 3}
	target3 := 6
	fmt.Printf("target:  %v numbers:  %v >>>  result:  %v \n", target3, nums3, twoSum(nums3, target3))
	fmt.Println("---------------------------")
}
