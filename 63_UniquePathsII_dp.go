package main

import "fmt"

func main() {
	//obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	obstacleGrid := [][]int{{0, 1}, {0, 0}}
	result := uniquePathsWithObstacles(obstacleGrid)
	fmt.Println(result)
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		if dp[i-1][0] > 0 && obstacleGrid[i][0] == 0 {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
		}
	}
	for j := 1; j < n; j++ {
		if dp[0][j-1] > 0 && obstacleGrid[0][j] == 0 {
			dp[0][j] = 1
		} else {
			dp[0][j] = 0
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = 0
			}
		}
	}
	return dp[m-1][n-1]
}
