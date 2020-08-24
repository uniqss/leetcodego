package main

import "fmt"

func main() {
	str := "MCMXCIV"
	ret := romanToInt(str)
	fmt.Println(ret)
}

func romanToInt(s string) int {
	ret := 0
	dict := map[uint8]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for i := 0; i < len(s); i++ {
		post := i + 1
		if post < len(s) && dict[s[i]] < dict[s[post]] {
			ret += dict[s[post]] - dict[s[i]]
			i++
			continue
		}
		ret += dict[s[i]]
	}

	return ret
}
