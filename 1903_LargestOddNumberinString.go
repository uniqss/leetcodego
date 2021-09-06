package main

import (
	"fmt"
)

func main() {
	s := "52"
	ret := largestOddNumber(s)
	fmt.Println(ret)
}

func largestOddNumber(num string) string {
	for i:= len(num) - 1;i >= 0;i-- {
		if num[i] % 2 == 1 {
			return num[0:i+1]
		}
	}
	return ""
}
