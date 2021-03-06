package main

import "fmt"

func main() {
	m := 3
	n := 7
	result := uniquePaths(m, n)
	fmt.Println(result)
}

func uniquePaths(m int, n int) int {
	if m > n {
		tmp := m
		m = n
		n = tmp
	}

	dp := make([]int, m)
	for i := 0; i < m; i++ {
		dp[i] = 1
	}
	left := 0
	for j := 1; j < n; j++ {
		left = 0
		for i := 0; i < m; i++ {
			dp[i] += left
			left = dp[i]
		}
	}
	return dp[m-1]
}
