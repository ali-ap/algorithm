package problems

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	sign := "+"
	r := ""
	str := strconv.Itoa(x)
	if str[0] == '-' {
		str = str[1:]
		sign = "-"
	}
	for i := 0; i < len(str); i++ {
		r = string(str[i]) + r
	}
	rx, err := strconv.ParseInt(sign+r, 10, 32)
	if err != nil {
		return 0
	}

	return int(rx)
}

func reverse2(x int) int {
	r := 0
	for x != 0 {
		l := x % 10
		x /= 10
		c := r*10 + l
		if c > math.MaxInt32 || c < math.MinInt32 {
			return 0
		}
		r = c
	}
	return r
}

func SolveReverseInteger() {
	fmt.Println("reverse integer")
	fmt.Println("https://leetcode.com/problems/reverse-integer/")
	fmt.Println("---------------------------")

	s1 := 123
	fmt.Printf("the reverse of %v is %v \n", s1, reverse(s1))
	fmt.Println("---------------------------")

	s2 := -123
	fmt.Printf("the reverse of %v is %v \n", s2, reverse2(s2))
	fmt.Println("---------------------------")

	s3 := 120
	fmt.Printf("the reverse of %v is %v \n", s3, reverse(s3))
	fmt.Println("---------------------------")

}
