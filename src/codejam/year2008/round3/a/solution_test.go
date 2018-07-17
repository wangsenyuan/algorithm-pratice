package main

import "testing"

func runSample(t *testing.T, n int, actions []string, count []int, expect int) {
	res := solve(n, actions, count)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, actions, count, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	actions := []string{"FFFR"}
	count := []int{4}
	expect := 0
	runSample(t, n, actions, count, expect)
}

func TestSample2(t *testing.T) {
	n := 9
	actions := []string{"F", "R", "F", "RFF", "LFF", "LFFFR", "F", "R", "F"}
	count := []int{6, 1, 4, 2, 1, 1, 2, 1, 5}
	expect := 4
	runSample(t, n, actions, count, expect)
}
