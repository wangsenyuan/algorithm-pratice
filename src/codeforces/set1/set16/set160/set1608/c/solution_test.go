package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect string) {
	res := solve(A, B)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4}
	B := []int{1, 2, 3, 4}
	expect := "0001"
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{11, 12, 20, 21}
	B := []int{44, 22, 11, 30}
	expect := "1111"
	runSample(t, A, B, expect)
}
