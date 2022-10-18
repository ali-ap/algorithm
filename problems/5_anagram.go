package problems

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	//another way to do this is with map array but be sure to check if the key count is equal
	sa := strings.Split(s, "")
	sort.Strings(sa)
	ta := strings.Split(t, "")
	sort.Strings(ta)
	return strings.Join(sa, "") == strings.Join(ta, "")
}

func SolveAnagram() {

	fmt.Println("anagram")
	fmt.Println("https://leetcode.com/problems/valid-anagram/")
	fmt.Println("---------------------------")

	s1, t1 := "anagram", "nagaram"
	fmt.Printf("are %v anagram of  %v ? %v \n", s1, t1, isAnagram(s1, t1))
	fmt.Println("---------------------------")

	s2, t2 := "rat", "car"
	fmt.Printf("are %v anagram of  %v ? %v \n", s1, t1, isAnagram(s2, t2))
	fmt.Println("---------------------------")

}
