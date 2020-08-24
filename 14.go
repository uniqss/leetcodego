package main

import "fmt"

func main() {
	strs := []string{"asdf", "asd", "asdfasdf"}
	ret := longestCommonPrefix(strs)
	fmt.Println(ret)
}


func longestCommonPrefix(strs []string) string {
	strsLen := len(strs)
	if strsLen == 0 {
		return ""
	}
	if strsLen == 1 {
		return strs[0]
	}
	end := 0
	str := strs[0]
	for ;end < len(str);end++ {
		for i:=0;i < strsLen;i++ {
			strCurr := strs[i]
			if end >= len(strCurr) {
				return strs[i]
			}
			if str[end] != strCurr[end] {
				return str[0:end]
			}
		}
	}
	return str
}
