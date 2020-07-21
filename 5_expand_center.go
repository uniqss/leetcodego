package main

import "fmt"

func main() {
	//s := "babad"
	s := "cbbd"
	ret := longestPalindrome(s)
	fmt.Println(ret)
}

func longestPalindrome(s string) string {
	slen := len(s)
	if slen == 0 {
		return s
	}
	middle := 0
	count := -1
	leni := 0
	leni1 := 0
	var i int
	for i=0;i < slen;i++{
		leni = expandAroundCenter(s, i, i)
		leni1 = expandAroundCenter(s, i, i + 1)
		if count < leni{
			count = leni
			middle = i
		}
		if count < leni1{
			count = leni1
			middle = i
		}
	}
	return s[middle - (count - 1)/2:middle - (count - 1)/2 + count]
}

func expandAroundCenter(s string, left int, right int) int {
	slen := len(s)
	for ;left>=0 && right < slen && s[left] == s[right]; {
		left--
		right++
	}
	return right - left - 1
}
