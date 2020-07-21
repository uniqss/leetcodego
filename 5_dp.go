package main

import "fmt"

func main() {
	s := "babad"
	ret := longestPalindrome(s)
	fmt.Println(ret)
}

func longestPalindrome(s string) string {
	begin := 0
	maxlen := 0
	slen := len(s)
	vii := make([][]int, slen)
	for i := 0; i < slen; i++ {
		vii[i] = make([]int, slen)
	}
	for len := 0; len < slen; len++ {
		for i := 0; i+len < slen; i++ {
			j := i + len
			if len == 0 {
				vii[i][j] = len + 1
			} else if len == 1 {
				if s[i] == s[j] {
					vii[i][j] = len + 1
				} else {
					vii[i][j] = 0
				}
			} else {
				if s[i] == s[j] && vii[i+1][j-1] > 0 {
					vii[i][j] = len + 1
				} else {
					vii[i][j] = 0
				}
			}
			if vii[i][j] > maxlen {
				begin = i
				maxlen = vii[i][j]
			}
		}
	}
	return s[begin : begin+maxlen]
}
