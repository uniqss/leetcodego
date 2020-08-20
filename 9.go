package main

import "fmt"

func main() {
	i := 12321
	ret := isPalindrome(i)
	fmt.Println(ret)
}

func isPalindrome(x int) bool {
	if x < 0 || x > 0 && x%10 == 0 {
		return false
	}
	xreverse := 0
	xmod := 0
	for x > xreverse {
		xmod = x % 10
		x /= 10
		xreverse *= 10
		xreverse += xmod
	}
	if x == xreverse || x == xreverse/10 {
		return true
	}

	return false
}
