package main

import "testing"

func runSample(t *testing.T, n, m int, F [][]int, expect string) {
	_, res := solve(n, m, F)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %s, but got %s", n, m, F, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 1
	F := [][]int{{1, 2}}
	expect := "AE"

	runSample(t, n, m, F, expect)

}

func TestSample2(t *testing.T) {
	n, m := 4, 3
	F := [][]int{{1, 2}, {2, 3}, {3, 1}}
	expect := "-1"

	runSample(t, n, m, F, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 2
	F := [][]int{{1, 3}, {2, 3}}
	expect := "AAE"

	runSample(t, n, m, F, expect)
}
