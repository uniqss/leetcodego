package main

import "fmt"

func main() {
	n := 3
	result := climbStairs(n)
	fmt.Println(result)
}

func climbStairs(n int) int {
	dpi1 := 1
	dpi2 := 1
	result := 1
	for i := 2; i <= n; i++ {
		result = dpi1 + dpi2
		dpi2 = dpi1
		dpi1 = result
	}
	return result
}
