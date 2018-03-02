package main

import "testing"

func runSample(t *testing.T, n int, g []int, s []int, kp int, ks int, expect int) {
	res := solve(n, g, s, kp, ks)

	if res != expect {
		t.Errorf("Sample %d %v %v %d %d, expect %d, but got %d", n, g, s, kp, ks, expect, res)
	}
}

func TestSample1(t *testing.T) {
	kp, ks := 2, 3
	n := 3
	g := []int{3, 1, 3}
	s := []int{7, 3, 4}
	expect := 2
	runSample(t, n, g, s, kp, ks, expect)
}

func TestSample2(t *testing.T) {
	kp, ks := 2, 3
	n := 3
	g := []int{3, 1, 3}
	s := []int{2, 7, 4}
	expect := 0
	runSample(t, n, g, s, kp, ks, expect)
}
