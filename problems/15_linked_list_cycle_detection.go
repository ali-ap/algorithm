package problems

import "fmt"

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	tortoise, hare := head, head
	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next
		if hare == tortoise {
			return true
		}
	}
	return false
}

func SolveLinkListCycleDetection() {
	fmt.Println("Linked List Cycle Detection")
	fmt.Println("https://leetcode.com/problems/linked-list-cycle/")
	fmt.Println("---------------------------")
}
