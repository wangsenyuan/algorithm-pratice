package main

import "testing"

func runSample(t *testing.T, A []int, expect string) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 5, 3, 4, 2}
	expect := "10111"
	runSample(t, A, expect)
}
func TestSample2(t *testing.T) {
	A := []int{1, 3, 2, 1}
	expect := "0001"
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 3, 3, 3, 2}
	expect := "00111"
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := "1111111111"
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{3, 3, 2}
	expect := "000"
	runSample(t, A, expect)
}
