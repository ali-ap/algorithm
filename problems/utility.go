package problems

import "math"

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Queue []interface{}

func (q *Queue) Enqueue(element interface{}) {
	*q = append(*q, element)
}
func (q *Queue) Dequeue() interface{} {
	data := *q
	if len(data) < 1 {
		return nil
	}
	element := data[0]
	*q = data[1:len(data)]
	return element
}

func NewQueue() *Queue {
	return &Queue{}
}

type Stack []interface{}

func (s *Stack) Push(element interface{}) {
	*s = append(*s, element)
}
func (s *Stack) Pop() interface{} {
	data := *s
	length := len(data)
	if length < 1 {
		return nil
	}
	element := data[length-1]
	*s = data[:length-1]
	return element
}
func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}
func NewStack() *Stack {
	return &Stack{}
}

func BinarySearch(nums []int, left, right, target int) int {
	for left <= right {
		mid := int(math.Floor(float64((left + right) / 2)))
		midVal := nums[mid]
		if midVal == target {
			return mid
		} else if midVal < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
