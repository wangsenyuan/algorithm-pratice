package main

import "testing"

func runSample(t *testing.T, A []int, x int, expect int) {
	res := solve(A, x)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{12}
	x := 2
	expect := 36
	runSample(t, A, x, expect)
}

func TestSample2(t *testing.T) {
	A := []int{4, 6, 8, 2}
	// 4, 6, 8, 2  => 6, 8, 2, 2, 2, => 8, 2, 2, 2, 3, 3
	//  2, 2, 2, 3, 3, 4, 4 => 2, 2, 3, 3, 4, 4, 1, 1
	// 4 + 6 + 8 + 2 = 20
	x := 2
	expect := 44
	runSample(t, A, x, expect)
}

func TestSample3(t *testing.T) {
	A := []int{4, 6, 8, 9}
	x := 2
	expect := 45
	runSample(t, A, x, expect)
}
