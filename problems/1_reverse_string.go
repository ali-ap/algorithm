package problems

import (
	"fmt"
)

func ReverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

}

func SolveReverseString() {
	fmt.Println("reverse string")
	fmt.Println("https://leetcode.com/problems/reverse-string/")

	s1 := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Printf("original string: %c \n", s1)
	ReverseString(s1)
	fmt.Printf("result string: %c \n", s1)
	fmt.Println("---------------------------")
	s2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	fmt.Printf("original string: %c \n", s2)
	ReverseString(s2)
	fmt.Printf("result string: %c \n", s2)
}
