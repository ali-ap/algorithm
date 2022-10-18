package problems

import (
	"fmt"
	"strings"
)

func capitalizeTitle(title string) string {
	arr := strings.Split(title, " ")
	for k, v := range arr {
		if len(v) > 2 {
			arr[k] = strings.ToUpper(string(arr[k][0])) + strings.ToLower(arr[k][1:])
		} else {
			arr[k] = strings.ToLower(arr[k])
		}
	}
	return strings.Join(arr, " ")
}

func SolveCapitalizeTitle() {

	fmt.Println("capitalizeTitle")
	fmt.Println("https://leetcode.com/problems/capitalize-the-title/")
	fmt.Println("---------------------------")

	s1 := "capiTalIze tHe titLe"
	fmt.Printf("the capitalized sentece of  %v is >>>   %v \n", s1, capitalizeTitle(s1))
	fmt.Println("---------------------------")

	s2 := "First leTTeR of EACH Word"
	fmt.Printf("the capitalized sentece of  %v is >>>   %v \n", s2, capitalizeTitle(s2))
	fmt.Println("---------------------------")

	s3 := "i lOve leetcode"
	fmt.Printf("the capitalized sentece of  %v is >>>   %v \n", s3, capitalizeTitle(s3))
	fmt.Println("---------------------------")

}
