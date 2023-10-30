package main

import "testing"

func runSample(t *testing.T, a []int, m int, expect int) {
	res := solve(a, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{-1, -2, -3, -4}
	m := 3
	expect := 1
	runSample(t, a, m, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 4}
	m := 3
	expect := 1
	runSample(t, a, m, expect)
}

func TestSample3(t *testing.T) {
	a := []int{-2, 3, -5, 1, -20}
	m := 5
	expect := 0
	runSample(t, a, m, expect)
}

func TestSample4(t *testing.T) {
	a := []int{-2, 3, -5, -5, -20}
	m := 2
	expect := 3
	runSample(t, a, m, expect)
}

func TestSample5(t *testing.T) {
	a := []int{345875723, -48, 384678321, -375635768, -35867853, -35863586, -358683842, -81725678, 38576, -357865873}
	m := 4
	expect := 4
	runSample(t, a, m, expect)
}

func TestSample6(t *testing.T) {
	a := []int{1}
	m := 1
	expect := 0
	runSample(t, a, m, expect)
}
