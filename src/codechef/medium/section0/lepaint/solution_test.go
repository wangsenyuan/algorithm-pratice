package main

import (
	"testing"
	"math"
)

func runSample(t *testing.T, N int, C int, K int, L []int, R []int, expect float64) {
	res := solve(N, C, K, L, R)
	if math.Abs(res-expect) > 1e-7 {
		t.Errorf("Sample %d %d %d %v %v, expect %f, but got %f", N, C, K, L, R, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, C, K := 4, 3, 4
	L := []int{1, 2, 3, 1}
	R := []int{2, 4, 3, 4}
	var expect float64 = 3.444444444
	runSample(t, N, C, K, L, R, expect)
}

func TestSample2(t *testing.T) {
	N, C, K := 7, 10, 7
	L := []int{7, 3, 1, 5, 4, 2, 3}
	R := []int{7, 4, 2, 5, 5, 7, 3}
	expect := 22.943125000
	runSample(t, N, C, K, L, R, expect)
}
