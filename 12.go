package main

import "fmt"

func main() {
	i := 1994
	ret := intToRoman(i)
	fmt.Println(ret)
}

func intToRoman(num int) string {
	arrInt := []int{1000, 500, 100, 50, 10, 5, 1}
	arrString := []string{"M", "D", "C", "L", "X", "V", "I"}

	var ret string
	for i := 0; i < len(arrInt); i++ {
		for num >= arrInt[i] {
			ret += arrString[i]
			num -= arrInt[i]
		}
		post := i + (i + 1) % 2 + 1
		if post < len(arrInt) && num >= (arrInt[i] - arrInt[post]) {
			ret += arrString[post]
			ret += arrString[i]
			num -= arrInt[i] - arrInt[post]
		}
	}
	return ret
}
