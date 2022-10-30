package problems

import "fmt"

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	start, currentNode := head, head
	currentPos := 1

	for currentPos < left {
		start = currentNode
		currentNode = currentNode.Next
		currentPos++
	}
	tail := currentNode
	var newNode *ListNode
	for currentPos >= left && currentPos <= right {
		next := currentNode.Next
		currentNode.Next = newNode
		newNode = currentNode
		currentNode = next
		currentPos++
	}

	start.Next = newNode
	tail.Next = currentNode

	if left > 1 {
		return head
	} else {
		return newNode
	}
}

func SolveReverseBetween() {
	fmt.Println("Reverse Between Linked List")
	fmt.Println("https://leetcode.com/problems/reverse-linked-list-ii/")
	fmt.Println("---------------------------")
}
