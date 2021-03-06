package main

import "fmt"

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	result := minPathSum(grid)
	fmt.Println(result)
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = grid[i][j]
		}
	}

	for i := 1; i < m; i++ {
		dp[i][0] += dp[i-1][0]
	}

	for j := 1; j < n; j++ {
		dp[0][j] += dp[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i-1][j] <= dp[i][j-1] {
				dp[i][j] += dp[i-1][j]
			} else {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
