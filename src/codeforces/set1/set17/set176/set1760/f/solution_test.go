package main

import "testing"

func runSample(t *testing.T, a []int, c int, d int, expect int) {
	res := solve(a, c, d)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2}
	c, d := 5, 4
	expect := 2
	runSample(t, a, c, d, expect)
}

func TestSample2(t *testing.T) {
	a := []int{100, 10}
	c, d := 20, 10
	expect := -2
	runSample(t, a, c, d, expect)
}

func TestSample3(t *testing.T) {
	a := []int{7, 2, 6}
	c, d := 100, 3
	expect := -1
	runSample(t, a, c, d, expect)
}

func TestSample4(t *testing.T) {
	a := []int{4, 5, 6, 7}
	c, d := 20, 3
	expect := 1
	runSample(t, a, c, d, expect)
}
func TestSample5(t *testing.T) {
	a := []int{8217734, 927368, 26389746, 627896974}
	c, d := 100000000000, 2022
	expect := 12
	runSample(t, a, c, d, expect)
}

func TestSample6(t *testing.T) {
	a := []int{5, 1}
	c, d := 20, 4
	expect := 0
	runSample(t, a, c, d, expect)
}

func TestSample7(t *testing.T) {
	a := []int{1, 5, 3}
	c, d := 18, 8
	expect := 4
	runSample(t, a, c, d, expect)
}

func TestSample8(t *testing.T) {
	a := []int{2, 4, 5, 4, 3}
	c, d := 25, 5
	expect := 0
	runSample(t, a, c, d, expect)
}
