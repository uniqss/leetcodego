package main

import (
	"fmt"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ret := maxArea(height)
	fmt.Println(ret)
}

func min(left, right int) int {
	if left < right {
		return left
	} else {
		return right
	}
}
func calcVolume(left, right int, height []int) int {
	return (right - left) * min(height[left], height[right])
}
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	temp := 0
	volume := calcVolume(left, right, height)
	for {
		if height[left] < height[right] {
			temp = left + 1
			for temp < right && height[temp] < height[left] {
				temp++
			}
			if temp >= right {
				break
			}
			newVolume := calcVolume(temp, right, height)
			if newVolume > volume {
				volume = newVolume
			}
			left = temp
		} else {
			temp = right - 1
			for temp > left && height[temp] < height[right] {
				temp--
			}
			if temp <= left {
				break
			}
			newVolume := calcVolume(left, temp, height)
			if newVolume > volume {
				volume = newVolume
			}
			right = temp
		}
	}
	return volume
}
