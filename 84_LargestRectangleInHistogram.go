package main

import "fmt"

func main() {
	//arr := []int{4, 2, 0, 3, 2, 5}
	arr := []int{2, 1, 5, 6, 2, 3}
	result := largestRectangleArea(arr)
	fmt.Println(result)
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	var maxArea int = 0
	heights = append(heights, 0)

	var stackIdx []int

	for i := 0; i < len(heights); i++ {
		for len(stackIdx) > 0 && heights[i] < heights[stackIdx[len(stackIdx)-1]] {
			idx := stackIdx[len(stackIdx)-1]
			stackIdx = stackIdx[:len(stackIdx)-1]
			width := i
			if len(stackIdx) > 0 {
				width = i - stackIdx[len(stackIdx)-1] - 1
			}
			area := width * heights[idx]
			maxArea = max(maxArea, area)
		}
		stackIdx = append(stackIdx, i)
	}
	return maxArea
}
