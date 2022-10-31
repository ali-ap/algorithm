package problems

import (
	"fmt"
	"strings"
)

func minRemoveToMakeValid(s string) string {
	stack := NewStack()
	res := strings.Split(s, "")
	for i := 0; i < len(res); i++ {
		if res[i] == "(" {
			stack.Push(i)
		} else if res[i] == ")" && !stack.isEmpty() {
			stack.Pop()
		} else if res[i] == ")" {
			res[i] = ""
		}
	}

	for !stack.isEmpty() {
		p := (stack.Pop()).(int)
		res[p] = ""
	}
	return strings.Join(res, "")
}

func SolveMinRemoveToMakeValid() {
	fmt.Println("1249. Minimum Remove to Make Valid Parentheses")
	fmt.Println("https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/")
	fmt.Println("---------------------------")

	s := "))(("
	fmt.Println(minRemoveToMakeValid(s))
}
