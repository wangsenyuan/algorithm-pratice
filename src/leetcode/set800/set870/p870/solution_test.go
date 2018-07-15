package main

import (
	"testing"
)

func countAdvantage(B, A []int) int {
	var cnt int
	for i := 0; i < len(A); i++ {
		if A[i] > B[i] {
			cnt++
		}
	}
	return cnt
}

func runSample(t *testing.T, A []int, B []int, expect []int) {
	res := advantageCount(A, B)
	cnt1 := countAdvantage(B, expect)
	cnt2 := countAdvantage(B, res)
	if cnt1 != cnt2 {
		t.Errorf("Sample %v %v, expect %v with count %d, but got %v with %d", A, B, expect, cnt1, res, cnt2)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 7, 11, 15}
	B := []int{1, 10, 4, 11}
	expect := []int{2, 11, 7, 15}
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{12, 24, 8, 32}
	B := []int{13, 25, 32, 11}
	expect := []int{24, 32, 8, 12}
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 0, 4, 1, 2}
	B := []int{1, 3, 0, 0, 2}
	expect := []int{2, 0, 2, 1, 4}
	runSample(t, A, B, expect)
}

func TestSample4(t *testing.T) {
	A := []int{15777, 7355, 6475, 15448, 18412}
	B := []int{986, 13574, 14234, 18412, 19893}
	expect := []int{6475, 15448, 15777, 18412, 7355}
	runSample(t, A, B, expect)
}

func TestSample5(t *testing.T) {
	A := []int{8, 2, 4, 4, 5, 6, 6, 0, 4, 7}
	B := []int{0, 8, 7, 4, 4, 2, 8, 5, 2, 0}
	expect := []int{4, 7, 8, 6, 5, 4, 0, 6, 4, 2}
	runSample(t, A, B, expect)
}
