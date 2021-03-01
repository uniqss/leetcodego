package main

import (
	"fmt"
)

func main() {
	s := "(()"
	ret := longestValidParentheses(s)
	fmt.Println(ret)
}

func maxInt(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func longestValidParentheses(s string) int {
	if s == "" {
		return 0
	}

	var dp []int
	for i := 0; i < len(s); i++ {
		dp = append(dp, 0)
	}

	ret := 0

	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i-2 >= 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				dp[i] = dp[i-1] + 2
				if i-dp[i] >= 0 {
					dp[i] += dp[i-dp[i]]
				}
			}
		}
		ret = maxInt(ret, dp[i])
	}

	return ret
}
