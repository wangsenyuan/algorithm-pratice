package main

import "testing"

func runSample(t *testing.T, n int, C []string, L, R []int, expect int) {
	res := solve(n, C, L, R)
	if res != expect {
		t.Errorf("Sample %d %v %v %v, expect %d, but got %d", n, C, L, R, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	C := []string{"BLUD", "RED"}
	L := []int{1, 5001}
	R := []int{5000, 10000}
	expect := 2
	runSample(t, n, C, L, R, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	C := []string{"BLUD", "RED", "WHITE"}
	L := []int{1, 2000, 7000}
	R := []int{6000, 8000, 10000}
	expect := 3
	runSample(t, n, C, L, R, expect)
}
func TestSample3(t *testing.T) {
	n := 4
	C := []string{"BLUD", "RED", "ORANGE", "GREEN"}
	L := []int{1, 2000, 4000, 7000}
	R := []int{3000, 5000, 8000, 10000}
	expect := -1
	runSample(t, n, C, L, R, expect)
}

func TestSample4(t *testing.T) {
	n := 2
	C := []string{"BLUD", "RED"}
	L := []int{1, 4002}
	R := []int{4000, 10000}
	expect := -1
	runSample(t, n, C, L, R, expect)
}

func TestSample5(t *testing.T) {
	n := 2
	C := []string{"BLUD", "RED", "ORANGE"}
	L := []int{1, 4000, 3000}
	R := []int{6000, 10000, 8000}
	expect := 2
	runSample(t, n, C, L, R, expect)
}
