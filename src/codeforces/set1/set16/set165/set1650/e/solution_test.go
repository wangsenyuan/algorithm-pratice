package main

import "testing"

func runSample(t *testing.T, A []int, d int, expect int) {
	res := solve(d, A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 5, 9}
	d := 12
	expect := 2
	runSample(t, A, d, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 5}
	d := 5
	expect := 1
	runSample(t, A, d, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 5}
	d := 5
	expect := 1
	runSample(t, A, d, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 2}
	d := 100
	expect := 1
	runSample(t, A, d, expect)
}

func TestSample5(t *testing.T) {
	A := []int{3, 6, 9, 12, 15}
	d := 15
	expect := 2
	runSample(t, A, d, expect)
}

func TestSample6(t *testing.T) {
	A := []int{1, 400000000, 500000000}
	d := 1000000000
	expect := 99999999
	runSample(t, A, d, expect)
}

func TestSample7(t *testing.T) {
	A := []int{3, 4}
	d := 10
	expect := 3
	runSample(t, A, d, expect)
}

func TestSample8(t *testing.T) {
	A := []int{1, 2}
	d := 2
	expect := 0
	runSample(t, A, d, expect)
}

func TestSample9(t *testing.T) {
	A := []int{6, 11, 12, 13}
	d := 15
	expect := 1
	runSample(t, A, d, expect)
}

func TestSample10(t *testing.T) {
	A := []int{17, 20}
	d := 20
	expect := 9
	runSample(t, A, d, expect)
}

func TestSample11(t *testing.T) {
	A := []int{1, 11, 12}
	d := 13
	expect := 0
	runSample(t, A, d, expect)
}
