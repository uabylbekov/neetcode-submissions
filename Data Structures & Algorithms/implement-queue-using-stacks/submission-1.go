type MyQueue struct {
	Stack1 []int
	Stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{Stack1: []int{}, Stack2: []int{}}
}

func (this *MyQueue) Push(x int) {
	this.Stack1 = append(this.Stack1, x)
}

func (this *MyQueue) Pop() int {
	if len(this.Stack2) == 0 {
		for len(this.Stack1) > 0 {
			last := this.Stack1[len(this.Stack1) - 1]
			this.Stack1 = this.Stack1[:len(this.Stack1) - 1]
			this.Stack2 = append(this.Stack2, last)
		}
	}
	last := this.Stack2[len(this.Stack2) - 1]
	this.Stack2 = this.Stack2[:len(this.Stack2) - 1]
	return last
}

func (this *MyQueue) Peek() int {
	if len(this.Stack2) == 0 {
		for len(this.Stack1) > 0 {
			last := this.Stack1[len(this.Stack1) - 1]
			this.Stack1 = this.Stack1[:len(this.Stack1) - 1]
			this.Stack2 = append(this.Stack2, last)
		}
	}
	return this.Stack2[len(this.Stack2) - 1]
}

func (this *MyQueue) Empty() bool {
	if len(this.Stack2) == 0 && len(this.Stack1) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
