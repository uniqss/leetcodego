package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ret := trap(arr)
	fmt.Println(ret)
}

func maxInt(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
func minInt(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	ret := 0
	size := len(height)

	var left_max []int
	left_max = append(left_max, height[0])
	for i := 1; i < size; i++ {
		left_max = append(left_max, maxInt(left_max[i-1], height[i]))
	}

	var right_max []int
	right_max = append(right_max, height[size-1])
	for i := 1; i < size; i++ {
		right_max = append(right_max, maxInt(right_max[i-1], height[size-i-1]))
	}
	for i := 0; i < size; i++ {
		ret += minInt(left_max[i], right_max[size-i-1]) - height[i]
	}

	return ret
}
