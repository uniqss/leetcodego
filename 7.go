package main

import (
	"fmt"
	"math"
)

func main() {
	i := 1534236469
	ret := reverse(i)
	fmt.Println(ret)
}

func reverse(x int) int {
	xmod := 0
	max1 := math.MaxInt32 / 10
	min1 := math.MinInt32 / 10
	ret := 0
	if x == 0 {
		return 0
	}
	for x != 0 {
		xmod = x % 10
		x = x / 10
		if ret > max1 || (ret >= max1 && xmod > 7) {
			return 0
		}
		if ret < min1 || (ret <= min1 && xmod < -8) {
			return 0
		}
		ret *= 10
		ret += xmod
	}
	return ret
}
