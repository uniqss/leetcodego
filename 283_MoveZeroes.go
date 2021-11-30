package main

import "fmt"

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5}
	moveZeroes(arr)
	fmt.Println(arr)
}

func moveZeroes(nums []int)  {
	j := 0
	for i:=0;i < len(nums);i++ {
		if nums[i] != 0 {
			if i != j {
				nums[j] = nums[i]
				nums[i] = 0
			}
			j++
		}
	}
}