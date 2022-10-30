package problems

import "fmt"

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}
	return prev
}

func SolveReverseLinkedList() {
	fmt.Println("Reverse Linked List")
	fmt.Println("https://leetcode.com/problems/reverse-linked-list/")
	fmt.Println("---------------------------")

}
