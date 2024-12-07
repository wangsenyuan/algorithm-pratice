package main

import "testing"

func runSample(t *testing.T, A []int, D []int, K int, expect int) {
	res := solve(K, A, D)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	K := 1
	A := []int{5, 6, 7, 8, 10, 2}
	D := []int{3, 5, 6, 7, 1, 10}
	expect := 35
	runSample(t, A, D, K, expect)
}

func TestSample2(t *testing.T) {
	K := 0
	A := []int{5, 3, 2, 2}
	D := []int{13, 8, 5, 1}
	expect := 1
	runSample(t, A, D, K, expect)
}

func TestSample3(t *testing.T) {
	K := 2
	A := []int{1, 2, 4, 8}
	D := []int{5, 1, 3, 5}
	expect := 14
	runSample(t, A, D, K, expect)
}

func TestSample4(t *testing.T) {
	K := 3
	A := []int{1, 2, 4, 8}
	D := []int{1, 5, 3, 5}
	expect := 14
	runSample(t, A, D, K, expect)
}

func TestSample5(t *testing.T) {
	K := 2
	A := []int{1, 2, 4, 8}
	D := []int{17, 15, 16, 1}
	expect := 7
	runSample(t, A, D, K, expect)
}

func TestSample6(t *testing.T) {
	K := 1
	A := []int{1, 10, 10, 5, 20, 100}
	D := []int{1000000, 1, 10, 10, 100, 1000000}
	expect := 144
	runSample(t, A, D, K, expect)
}

func TestSample7(t *testing.T) {
	K := 1
	A := []int{1, 2, 2, 2, 1}
	D := []int{1, 101, 101, 101, 100}
	expect := 6
	runSample(t, A, D, K, expect)
}
