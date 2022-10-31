package problems

import "fmt"

type MyQueue struct {
	in  *Stack
	out *Stack
}

func Constructor() MyQueue {
	return MyQueue{
		in:  NewStack(),
		out: NewStack(),
	}
}

func (this *MyQueue) Push(x int) {
	*this.in = append(*this.in, x)
}

func (this *MyQueue) Pop() int {
	if this.out.isEmpty() {
		for !this.in.isEmpty() {
			this.out.Push(this.in.Pop())
		}
	}
	return this.out.Pop().(int)
}

func (this *MyQueue) Peek() int {
	if this.out.isEmpty() {
		for !this.in.isEmpty() {
			this.out.Push(this.in.Pop())
		}
	}
	data := *this.out

	return data[len(*this.out)-1].(int)
}

func (this *MyQueue) Empty() bool {
	return len(*this.in) == 0 && len(*this.out) == 0
}

func SolveMyQueue() {
	fmt.Println("232. Implement Queue using Stacks")
	fmt.Println("https://leetcode.com/problems/implement-queue-using-stacks/")
	fmt.Println("---------------------------")
	q := Constructor()
	q.Push(2) // queue is: [1, 2] (leftmost is front of the queue)
	q.Peek()  // return 1
	q.Pop()   // return 1, queue is [2]
	q.Empty()
	q.Push(1) // queue is: [1]
}
