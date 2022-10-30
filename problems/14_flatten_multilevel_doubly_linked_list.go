package problems

import (
	"fmt"
)

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}
	currentNode := root

	for currentNode != nil {
		if currentNode.Child == nil {
			currentNode = currentNode.Next
		} else {
			tail := currentNode.Child

			for {
				if tail.Next == nil {
					break
				}
				tail = tail.Next
			}
			tail.Next = currentNode.Next
			if tail.Next != nil {
				tail.Next.Prev = tail
			}
			currentNode.Next = currentNode.Child
			currentNode.Next.Prev = currentNode
			currentNode.Child = nil
		}
	}
	return root
}

func SolveFlatten() {
	fmt.Println("Reverse Between Linked List")
	fmt.Println("https://leetcode.com/problems/reverse-linked-list-ii/")
	fmt.Println("---------------------------")
}
