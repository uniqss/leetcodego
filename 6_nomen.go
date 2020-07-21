package main

import "fmt"

func main() {
	//s := "babad"
	s := "PAYPALISHIRING"
	ret := convert(s, 3)
	fmt.Println(ret)
}

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	ret := ""
	sLen := len(s)
	wrapLen := numRows*2 - 2
	for l := 0; l < numRows; l++ {
		for down := l; down < sLen; down += wrapLen {
			ret += s[down : down+1]
			downmod := down % wrapLen
			if downmod > 0 && downmod < numRows-1 {
				n := numRows - 1
				up := 2 * n - downmod + down - downmod
				if up < sLen {
					ret += s[up : up + 1]
				}
			}
		}
	}
	return ret
}
