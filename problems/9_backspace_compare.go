package problems

import "fmt"

func backspaceCompare(s string, t string) bool {
	return trim(s) == trim(t)
}

func trim(s string) string {
	var stack []uint8
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

func CompareBackspaces() {

	fmt.Println("Backspace String Compare")
	fmt.Println("https://leetcode.com/problems/backspace-string-compare/")
	fmt.Println("---------------------------")

	s, t := "ab#c", "ad#c"
	fmt.Printf(" are:  %v and %v equal? >>>  result:  %t \n", s, t, backspaceCompare(s, t))
	fmt.Println("---------------------------")

	s, t = "ab##", "c#d#"
	fmt.Printf(" are:  %v and %v equal? >>>  result:  %t \n", s, t, backspaceCompare(s, t))
	fmt.Println("---------------------------")

	s, t = "a#c", "b"
	fmt.Printf(" are:  %v and %v equal? >>>  result:  %t \n", s, t, backspaceCompare(s, t))
	fmt.Println("---------------------------")

}
