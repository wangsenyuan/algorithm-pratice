package main

import "testing"

func runSample(t *testing.T, p []int, expect int) {
	res := solve(p)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{0}
	expect := 1
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := []int{1, 0}
	expect := 2
	runSample(t, p, expect)
}

func TestSample3(t *testing.T) {
	p := []int{1, 0, 2}
	expect := 4
	runSample(t, p, expect)
}

func TestSample4(t *testing.T) {
	p := []int{0, 2, 1, 3}
	expect := 4
	runSample(t, p, expect)
}

func TestSample5(t *testing.T) {
	p := []int{3, 1, 0, 2, 4}
	expect := 8
	runSample(t, p, expect)
}

func TestSample6(t *testing.T) {
	p := []int{2, 0, 4, 1, 3, 5}
	expect := 8
	runSample(t, p, expect)
}

func TestSample7(t *testing.T) {
	p := []int{3, 7, 2, 6, 0, 1, 5, 4}
	expect := 15
	runSample(t, p, expect)
}

func TestSample8(t *testing.T) {
	p := []int{2, 0, 1, 3}
	expect := 6
	runSample(t, p, expect)
}
