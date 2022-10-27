package problems

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	var seen = make(map[byte]int)
	left, longest := 0, 0
	for right := 0; right < len(s); right++ {
		v, exist := seen[s[right]]
		if exist && v >= left {
			left = v + 1
		}
		seen[s[right]] = right
		longest = Max(longest, right-left+1)
	}
	return longest
}

func SolveLengthOfLongestSubstring() {

	fmt.Println("Length of longest substring")
	fmt.Println("https://leetcode.com/problems/longest-substring-without-repeating-characters/")
	fmt.Println("---------------------------")

	s := "abcabcbb"
	fmt.Printf(" the longest substring in:  %v is >>>  result:  %v \n", s, lengthOfLongestSubstring(s))
	fmt.Println("---------------------------")

	s = "bbbbb"
	fmt.Printf(" the longest substring in:  %v is >>>  result:  %v \n", s, lengthOfLongestSubstring(s))
	fmt.Println("---------------------------")

	s = "pwwkew"
	fmt.Printf(" the longest substring in:  %v is >>>  result:  %v \n", s, lengthOfLongestSubstring(s))
	fmt.Println("---------------------------")

}
