package main

import "fmt"

func main() {
	word1 := "horse"
	word2 := "ros"
	result := minDistance(word1, word2)
	fmt.Println(result)
}

func minInt(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	if len1 == 0 {
		return len2
	}
	if len2 == 0 {
		return len1
	}

	var dp [][]int
	for i := 0; i <= len1; i++ {
		var line []int
		for j := 0; j <= len2; j++ {
			line = append(line, i+j)
		}
		dp = append(dp, line)
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				disReplace := dp[i-1][j-1]
				disInsert := dp[i][j-1]
				disDelete := dp[i-1][j]
				dp[i][j] = minInt(disReplace, minInt(disInsert, disDelete)) + 1
			}
		}
	}
	return dp[len1][len2]
}
