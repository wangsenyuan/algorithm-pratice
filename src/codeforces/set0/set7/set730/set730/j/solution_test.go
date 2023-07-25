package main

import "testing"

func runSample(t *testing.T, A, B []int, cnt int, move int) {
	a, b := solve(A, B)

	if a != cnt || b != move {
		t.Fatalf("Sample expect %d %d, but got %d %d", cnt, move, a, b)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 3, 4, 3}
	B := []int{4, 7, 6, 5}
	a, b := 2, 6
	runSample(t, A, B, a, b)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1}
	B := []int{100, 100}
	a, b := 1, 1
	runSample(t, A, B, a, b)
}

func TestSample3(t *testing.T) {
	A := []int{10, 30, 5, 6, 24}
	B := []int{10, 41, 7, 8, 24}
	a, b := 3, 11
	runSample(t, A, B, a, b)
}

func TestSample4(t *testing.T) {
	A := []int{24, 22, 4, 34, 76, 13, 78, 1, 81, 51, 72, 11, 25, 46, 22, 33, 60, 42, 25, 19}
	B := []int{40, 81, 10, 34, 84, 16, 90, 38, 99, 81, 100, 19, 79, 65, 26, 80, 62, 47, 76, 47}
	a, b := 9, 217
	runSample(t, A, B, a, b)
}

func TestSample5(t *testing.T) {
	A := []int{18, 42, 5, 1, 26, 8, 40, 34, 8, 29}
	B := []int{18, 71, 21, 67, 38, 13, 99, 37, 47, 76}
	a, b := 3, 100
	runSample(t, A, B, a, b)
}
