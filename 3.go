package main

import (
	"fmt"
	"math"
)

func main() {
	s := "asdf"
	longest := lengthOfLongestSubstring(s)
	fmt.Println(longest)
}

func lengthOfLongestSubstring(s string) int {
	i := 0
	j := 0
	ret := 0
	arr := [256]int{}
	for tmp := range arr {
		arr[tmp] = -1
	}
	slen := len(s)
	for ; i < slen; i++ {
		for j = i; j < slen; j++ {
			if arr[s[j]] >= i {
				ret = int(math.Max(float64(ret), float64(j-i)))
				i = arr[s[j]]
				for tmp := range arr {
					arr[tmp] = -1
				}
				break
			}
			arr[s[j]] = j
		}
		if j == slen{
			ret = int(math.Max(float64(ret), float64(j-i)))
			break
		}
	}
	return ret
}
