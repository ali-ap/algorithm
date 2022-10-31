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
func (q *Queue) isEmpty() bool {
	return len(*q) == 0
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) BreadthFirstSearch() []int {
	var result []int
	queue := NewQueue()
	queue.Enqueue(*t)
	for !queue.isEmpty() {
		currentNode := queue.Dequeue().(TreeNode)
		result = append(result, currentNode.Val)
		if currentNode.Left != nil {
			queue.Enqueue(currentNode.Left)
		}
		if currentNode.Right != nil {
			queue.Enqueue(currentNode.Right)
		}
	}
	return result
}

func (t *TreeNode) DepthFirstInOrderSearch() (result []int) {
	t.dfsInOrder(result)
	return
}

func (t *TreeNode) DepthFirstPreOrderSearch() (result []int) {
	t.dfsPreOrder(result)
	return
}

func (t *TreeNode) DepthFirstPostOrderSearch() (result []int) {
	t.dfsPostOrder(result)
	return
}

func (t *TreeNode) dfsInOrder(result []int) {
	if t.Left != nil {
		t.Left.dfsInOrder(result)
	}
	result = append(result, t.Val)
	if t.Right != nil {
		t.Right.dfsInOrder(result)
	}
}

func (t *TreeNode) dfsPreOrder(result []int) {
	result = append(result, t.Val)
	if t.Left != nil {
		t.Left.dfsInOrder(result)
	}
	if t.Right != nil {
		t.Right.dfsInOrder(result)
	}
}

func (t *TreeNode) dfsPostOrder(result []int) {
	if t.Left != nil {
		t.Left.dfsInOrder(result)
	}
	if t.Right != nil {
		t.Right.dfsInOrder(result)
	}
	result = append(result, t.Val)
}

//max heap and min heap and priority queue
