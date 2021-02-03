package leetcode

type MinStack struct {
	slice []int
	min   int
}

/** initialize your data structure here. */
func Constructor() (this MinStack) {
	return
}

func (this *MinStack) Push(x int) {
	if len(this.slice) == 0 {
		this.min = x
	} else if this.min > x {
		this.min = x
	}
	this.slice = append(this.slice, x)
}

func (this *MinStack) Pop() {
	this.slice = this.slice[:len(this.slice)-1]
	if len(this.slice) == 1 {
		this.min = this.slice[0]
	} else if len(this.slice) > 1 {
		newmin := this.slice[0]
		for i := 1; i < len(this.slice); i++ {
			if this.slice[i] < newmin {
				newmin = this.slice[i]
			}
		}
		this.min = newmin
	}
}

func (this *MinStack) Top() int {
	if len(this.slice) < 1 {
		return 0
	}
	val := this.slice[len(this.slice)-1]
	return val
}

func (this *MinStack) GetMin() int {

	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
