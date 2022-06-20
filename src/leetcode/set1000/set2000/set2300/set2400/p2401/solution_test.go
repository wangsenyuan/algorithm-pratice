package p2401

import "testing"

func runSample(t *testing.T, m, n int, prices [][]int, expect int) {
	res := int(sellingWood(m, n, prices))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 3
	n := 5
	prices := [][]int{{1, 4, 2}, {2, 2, 7}, {2, 1, 3}}
	expect := 19
	runSample(t, m, n, prices, expect)
}

func TestSample2(t *testing.T) {
	m := 4
	n := 6
	prices := [][]int{{3, 2, 10}, {1, 4, 2}, {4, 1, 3}}
	expect := 32
	runSample(t, m, n, prices, expect)
}
