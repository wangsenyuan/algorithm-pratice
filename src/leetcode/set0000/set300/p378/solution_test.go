package main

import "testing"

func runSample(t *testing.T, m [][]int, k int, expect int) {
	res := kthSmallest(m, k)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", m, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := [][]int{{-1}}
	k := 1
	expect := -1
	runSample(t, m, k, expect)
}
