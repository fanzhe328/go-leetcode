package main

import "fmt"

func main() {
	cost := []int{0, 0, 0, 1}
	fmt.Println(minCostClimbingStairs(cost))

}

func minCostClimbingStairs(cost []int) int {
	c1 := cost[0]
	c2 := cost[1]
	cur := 0
	for i := 2; i < len(cost); i++ {
		cur = min(c1, c2) + cost[i]
		c1 = c2
		c2 = cur
	}
	return min(c1, c2)
}

// 10, 0, 10
// 15, 10, 15
// 30, 15, 30
// min(15, 30)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 1: cost[0]
// 2: min(cost[0], cost[1])
// 3: min(cost[1], cost[2])
