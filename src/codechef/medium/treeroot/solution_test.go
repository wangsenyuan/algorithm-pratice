package main

import (
	"testing"
)

func runSample(t *testing.T, n int, node []int, sum []int, expect int) {
	res := solve(n, node, sum)
	if res != expect {
		t.Errorf("sample %d %v, %v, expect %d, but got %d", n, node, sum, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	node := []int{4}
	sum := []int{0}
	runSample(t, n, node, sum, 4)
}
