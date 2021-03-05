package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ret := maxSubArray(nums)
	fmt.Println(ret)
}

func maxInt(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var dp = make([]int, len(nums))
	dp[0] = nums[0]
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = maxInt(dp[i-1]+nums[i], nums[i])
		ret = maxInt(ret, dp[i])
	}
	return ret
}
