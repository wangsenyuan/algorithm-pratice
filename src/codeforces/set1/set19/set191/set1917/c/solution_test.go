package main

import "testing"

func runSample(t *testing.T, a []int, d int, v []int, expect int) {
	res := solve(a, d, v)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	d := 4
	a := []int{1, 2, 3}
	v := []int{1, 3, 2, 3}
	expect := 4
	runSample(t, a, d, v, expect)
}

func TestSample2(t *testing.T) {
	d := 3
	a := []int{6, 1, 2, 4, 1, 5}
	v := []int{6, 6}
	expect := 3
	runSample(t, a, d, v, expect)
}

func TestSample3(t *testing.T) {
	d := 1
	a := []int{0, 5, 0, 5, 0}
	v := []int{5}
	expect := 0
	runSample(t, a, d, v, expect)
}

func TestSample4(t *testing.T) {
	d := 1
	a := []int{1}
	v := []int{1}
	expect := 1
	runSample(t, a, d, v, expect)
}
func TestSample5(t *testing.T) {
	d := 6
	a := []int{1, 2, 3}
	v := []int{1, 3, 2, 3}
	expect := 5
	runSample(t, a, d, v, expect)
}
