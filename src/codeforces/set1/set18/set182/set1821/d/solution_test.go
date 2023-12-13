package main

import "testing"

func runSample(t *testing.T, k int, L, R []int, expect int) {
	res := solve(k, L, R)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 20
	L := []int{10, 13, 16, 19}
	R := []int{11, 14, 17, 20}
	expect := -1
	runSample(t, k, L, R, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	L := []int{1, 3}
	R := []int{1, 4}
	expect := 8
	runSample(t, k, L, R, expect)
}

func TestSample3(t *testing.T) {
	k := 3
	L := []int{1, 3}
	R := []int{1, 10}
	expect := 7
	runSample(t, k, L, R, expect)
}

func TestSample4(t *testing.T) {
	k := 4
	L := []int{99, 999999999}
	R := []int{100, 1000000000}
	expect := 1000000004
	runSample(t, k, L, R, expect)
}

func TestSample5(t *testing.T) {
	k := 4
	L := []int{1, 5}
	R := []int{3, 10}
	// 1,2,3, 5
	// 4, 5, 6, 7
	expect := 9
	runSample(t, k, L, R, expect)
}

func TestSample6(t *testing.T) {
	k := 3
	L := []int{1, 3, 5}
	R := []int{1, 3, 5}
	// 1,2,3, 5
	// 4, 5, 6, 7
	expect := 11
	runSample(t, k, L, R, expect)
}
