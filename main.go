package main

import "fmt"

// FindTargetSumWays counts the ways to put + or - before every number so
// the expression equals target. Algebra reduces it to subset counting:
// the positives must sum to (total + target) / 2.
func FindTargetSumWays(nums []int, target int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	if target > total || target < -total || (total+target)%2 != 0 {
		return 0
	}
	goal := (total + target) / 2
	dp := make([]int, goal+1)
	dp[0] = 1
	for _, v := range nums {
		for s := goal; s >= v; s-- {
			dp[s] += dp[s-v]
		}
	}
	return dp[goal]
}

func main() {
	fmt.Println(FindTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(FindTargetSumWays([]int{1}, 2))
}
