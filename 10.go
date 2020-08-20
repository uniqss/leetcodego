package main

import "fmt"

func main() {
	s := "aaaabbbb"
	p := "c*a*b*"
	match := isMatch(s, p)
	fmt.Println(match)
}

func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)

	var f [][]bool = make([][]bool, sLen+1)
	for j := 0; j < sLen+1; j++ {
		f[j] = make([]bool, pLen+1)
	}

	f[0][0] = true

	for j := 2; j <= pLen; j++ {
		f[0][j] = f[0][j-2] && p[j-1] == '*'
	}

	for i := 1; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if p[j-1] != '*' {
				f[i][j] = f[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			} else {
				f[i][j] = f[i][j-2] || (f[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			}
		}
	}

	return f[sLen][pLen]
}
