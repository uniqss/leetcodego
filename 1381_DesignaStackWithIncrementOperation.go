package main

import (
	"fmt"
)

func main() {
	cmds := []string{"CustomStack","push","push","pop","push","push","push","increment","increment","pop","pop","pop","pop"}
	params := [][]int {{3},{1},{2},{},{2},{3},{4},{5,100},{2,100},{},{},{},{}}

	var obj CustomStack
	for i:=0;i< len(cmds);i++ {
		cmd := cmds[i]
		param := params[i]
		switch cmd {
		case "CustomStack":
			obj = Constructor(param[0])
		case "push":
			obj.Push(param[0])
		case "pop":
			obj.Pop()
		case "increment":
			obj.Increment(param[0], param[1])
		}
		fmt.Println(obj)
	}
}


type CustomStack struct {
	elements []int
	maxSize int
	nextIdx int
}

func Constructor(maxSize int) CustomStack {
	ret := CustomStack{}
	ret.maxSize = maxSize
	ret.nextIdx = 0
	for i:=0;i < maxSize;i++ {
		ret.elements = append(ret.elements, 0)
	}
	return ret
}

func (this *CustomStack) Push(x int)  {
	if this.nextIdx >= this.maxSize {
		return
	}
	this.elements[this.nextIdx] = x
	this.nextIdx++
}

func (this *CustomStack) Pop() int {
	if this.nextIdx == 0 {
		return -1
	}
	this.nextIdx--
	ret := this.elements[this.nextIdx]
	this.elements[this.nextIdx] = 0
	return ret
}

func (this *CustomStack) Increment(k int, val int)  {
	for i:=0;i < this.nextIdx && i < k;i++ {
		this.elements[i] += val
	}
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */