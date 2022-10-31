package problems

import "fmt"

func isValid(s string) bool {
	parentheses := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		_, ok := parentheses[s[i]]
		if ok {
			stack.Push(s[i])
		} else {
			if stack.isEmpty() {
				return false
			}
			pair := stack.Pop()
			correctPair, ok := parentheses[pair.(byte)]
			if !ok || s[i] != correctPair {
				return false
			}
		}
	}
	return len(*stack) == 0
}

func SolveValidParentheses() {
	fmt.Println("Valid Parentheses")
	fmt.Println("https://leetcode.com/problems/valid-parentheses/submissions/")
	fmt.Println("---------------------------")

	s := "()"

	println(isValid(s))

}
