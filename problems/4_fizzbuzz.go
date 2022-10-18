package problems

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	var r []string
	for i := 1; i <= n; i++ {
		txt := ""
		if i%3 == 0 {
			txt += "Fizz"
		}
		if i%5 == 0 {
			txt += "Buzz"
		}
		if txt == "" {
			txt = strconv.Itoa(i)
		}
		r = append(r, txt)
	}
	return r
}

func SolveFizzBuz() {
	fmt.Println("fizzbuzz")
	fmt.Println("https://leetcode.com/problems/fizz-buzz/")
	fmt.Println("---------------------------")

	s1 := 3
	fmt.Printf("the fizzbuzz of %v is %v \n", s1, fizzBuzz(s1))
	fmt.Println("---------------------------")

	s2 := 5
	fmt.Printf("the fizzbuzz of %v is %v \n", s2, fizzBuzz(s2))
	fmt.Println("---------------------------")

	s3 := 15
	fmt.Printf("the fizzbuzz of %v is %v \n", s3, fizzBuzz(s3))
	fmt.Println("---------------------------")

}
