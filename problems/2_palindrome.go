package problems

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {

	//we can also create a reverse string a just compare it with the actual string

	s = strings.ToLower(regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllLiteralString(s, ""))
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func IsValidPalindrome() {
	fmt.Println("palindrome string")
	fmt.Println("https://leetcode.com/problems/valid-palindrome/")
	fmt.Println("---------------------------")
	s1 := "A man, a plan, a canal: Panama"
	fmt.Printf("Is string \" %v \" a valid palindrome? %v \n", s1, isPalindrome(s1))
	fmt.Println("---------------------------")
	s2 := "race a car"
	fmt.Printf("Is string \" %v  \" a valid palindrome? %v \n", s2, isPalindrome(s2))
	fmt.Println("---------------------------")
	s3 := " "
	fmt.Printf("Is string \" %v  \" a valid palindrome? %v \n", s3, isPalindrome(s3))
}
