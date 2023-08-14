package main

import "testing"

func runSample(t *testing.T, n int, k int, A []int, expect string) {
	res := solve(n, k, A)
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 13
	A := []int{0, 1, 2}
	expect := "59"
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 2, 777
	A := []int{0, 4}
	expect := "778"
	runSample(t, n, k, A, expect)
}

func TestSample3(t *testing.T) {
	n, k := 3, 255
	A := []int{0, 1, 3}
	expect := "148999"
	runSample(t, n, k, A, expect)
}

func TestSample4(t *testing.T) {
	n, k := 10, 1000000000
	A := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := "999999920999999999"
	runSample(t, n, k, A, expect)
}
