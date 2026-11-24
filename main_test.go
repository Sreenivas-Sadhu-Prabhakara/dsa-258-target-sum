package main

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 1, 1, 1, 1}, 3, 5},
		{[]int{1}, 1, 1},
		{[]int{1}, 2, 0},
		{[]int{1, 2, 1}, 0, 2}, // +1+2... check: +1-2+1=0, -1+2-1=0 -> 2
		{[]int{0, 0, 1}, 1, 4}, // zeros take either sign freely
		{[]int{100}, -100, 1},  // negative targets by symmetry
		{[]int{2, 2}, 1, 0},    // parity mismatch
	}
	for _, c := range cases {
		if got := FindTargetSumWays(c.nums, c.target); got != c.want {
			t.Errorf("FindTargetSumWays(%v,%d) = %d, want %d", c.nums, c.target, got, c.want)
		}
	}
}
