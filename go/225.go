package leetcode

type MyStack struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor() (this MyStack) {
	return
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	lqueue := len(this.queue)
	this.queue = append(this.queue, x)
	for ; lqueue > 0; lqueue-- {
		this.queue = append(this.queue, this.queue[0])
		this.queue = this.queue[1:]
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	val := this.queue[0]
	this.queue = this.queue[1:]
	return val
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if len(this.queue) < 1 {
		return 0
	}
	val := this.queue[0]
	return val
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
