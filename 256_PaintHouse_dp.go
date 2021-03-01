package main

import (
	"fmt"
)

func minInt(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	var costs [][]int
	for i := 0; i < 11; i++ {
		var tmp []int
		tmp = append(tmp, 3)
		tmp = append(tmp, 5)
		tmp = append(tmp, 7)
		costs = append(costs, tmp)
	}
	ret := minCost2(costs)
	fmt.Println(ret)
}

func minCost(costs [][]int) int {
	i := 1
	for ; i < len(costs); i++ {
		costs[i][0] += minInt(costs[i-1][1], costs[i-1][2])
		costs[i][1] += minInt(costs[i-1][0], costs[i-1][2])
		costs[i][2] += minInt(costs[i-1][0], costs[i-1][1])
	}

	i--
	return minInt(minInt(costs[i][0], costs[i][1]), costs[i][2])
}

func minCost2(costs [][]int) int {
	current := [3]int{costs[0][0], costs[0][1], costs[0][2]}
	last := [3]int{costs[0][0], costs[0][1], costs[0][2]}
	for i := 1; i < len(costs); i++ {
		current[0] = costs[i][0] + minInt(last[1], last[2])
		current[1] = costs[i][1] + minInt(last[0], last[2])
		current[2] = costs[i][2] + minInt(last[0], last[1])

		for j := 0; j < 3; j++ {
			last[j] = current[j]
		}
	}

	return minInt(minInt(current[0], current[1]), current[2])
}
