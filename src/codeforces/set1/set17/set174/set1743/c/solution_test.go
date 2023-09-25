package main

import "testing"

func runSample(t *testing.T, s string, a []int, expect int) {
	res := solve(s, a)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "01110"
	a := []int{10, 5, 8, 9, 6}
	expect := 27
	runSample(t, s, a, expect)
}

func TestSample2(t *testing.T) {
	s := "011011"
	a := []int{20, 10, 9, 30, 20, 19}
	expect := 80
	runSample(t, s, a, expect)
}

func TestSample3(t *testing.T) {
	s := "0111"
	a := []int{5, 4, 5, 1}
	expect := 14
	runSample(t, s, a, expect)
}
