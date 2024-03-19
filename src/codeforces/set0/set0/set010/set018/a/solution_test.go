package main

import "testing"

func runSample(t *testing.T, pts []int, expect string) {
	res := solve(pts)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	pts := []int{0, 0, 1, 0, 4, 1}
	expect := "NEITHER"
	runSample(t, pts, expect)
}
