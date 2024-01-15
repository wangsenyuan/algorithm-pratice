package main

import "testing"

func runSample(t *testing.T, s []int, c int, expect int) {
	res := solve(s, c)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := []int{3, 2, 1}
	c := 50
	expect := 1
	runSample(t, s, c, expect)
}

func TestSample2(t *testing.T) {
	s := []int{6}
	c := 100
	expect := 2
	runSample(t, s, c, expect)
}

func TestSample3(t *testing.T) {
	s := []int{2609, 10, 5098, 9591, 949, 8485, 6385, 4586, 1064, 5412, 6564, 8460, 2245, 6552, 5089, 8353, 3803, 3764}
	c := 886531871815571953
	expect := 110961227
	runSample(t, s, c, expect)
}
