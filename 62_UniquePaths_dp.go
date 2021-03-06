package main

import "fmt"

func main() {
	m := 3
	n := 7
	result := uniquePaths(m, n)
	fmt.Println(result)
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	//for i := 0; i < m; i++ {
	//	dp[i] = make([]int, n)
	//}
	for i := 0; i < m; i++ {
		tmp := make([]int, n)
		for j := 0; j < n; j++ {
			tmp[j] = 1
		}
		dp[i] = tmp
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
