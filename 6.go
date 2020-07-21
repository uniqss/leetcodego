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
	arrLine := make([]string, numRows)
	goingDown := true
	lineIdx := 0
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		arrLine[lineIdx] += s[i : i+1]
		if goingDown {
			lineIdx++
		} else {
			lineIdx--
		}
		if lineIdx >= numRows-1 {
			goingDown = false
		}
		if lineIdx <= 0 {
			goingDown = true
		}
	}
	ret := ""
	for i := 0; i < numRows; i++ {
		ret += arrLine[i]
	}
	return ret
}
