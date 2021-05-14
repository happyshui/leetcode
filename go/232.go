package leetcode

type MyQueue struct {
	instack  []int
	outstack []int
}

/** Initialize your data structure here. */
func Constructor() (this MyQueue) {
	return
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.instack = append(this.instack, x)
	lqueue := len(this.instack)
	for ; lqueue > 0; lqueue-- {
		val := this.instack[lqueue-1]
		this.outstack = append(this.outstack, val)
		this.instack = this.instack[1:]
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	val := this.outstack[0]
	this.outstack = this.outstack[1:]
	return val
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	val := this.outstack[0]
	return val
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.outstack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
