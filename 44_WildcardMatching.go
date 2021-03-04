package main

import "fmt"

func main() {
	s := "abc"
	p := "*bc"
	result := isMatch(s, p)
	fmt.Println(result)
}

func isMatch(s string, p string) bool {
	slen := len(s)
	plen := len(p)
	var dp [][]bool
	for i := 0; i <= slen; i++ {
		var tmp []bool
		for j := 0; j <= plen; j++ {
			tmp = append(tmp, false)
		}
		dp = append(dp, tmp)
	}

	dp[0][0] = true
	for j := 1; j <= plen; j++ {
		for i := 0; i <= slen; i++ {
			if p[j-1] == '*' {
				if dp[i][j-1] {
					dp[i][j] = true
				}
				if i > 0 && dp[i-1][j] {
					dp[i][j] = true
				}
			} else if p[j-1] == '?' {
				if i > 0 && dp[i-1][j-1] {
					dp[i][j] = true
				}
			} else {
				if i > 0 && dp[i-1][j-1] && p[j-1] == s[i-1] {
					dp[i][j] = true
				}
			}
		}
	}

	return dp[slen][plen]
}
