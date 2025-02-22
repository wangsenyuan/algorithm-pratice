package main

import "testing"

func runSample(t *testing.T, u int, r int, A []int, B []int, K []int, P []int, expect int64) {
	res := solve(u, r, A, B, K, P)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	u, r := 2, 1
	A := []int{7, 7, 7}
	B := []int{8, 8, 8}
	K := []int{1, 2, 3}
	P := []int{1, 3, 2}
	var expect int64 = 96
	runSample(t, u, r, A, B, K, P, expect)
}

func TestSample2(t *testing.T) {
	u, r := 1, 0
	A := []int{1, 1}
	B := []int{1, 1}
	K := []int{1, -1}
	P := []int{1, 2}
	var expect int64
	runSample(t, u, r, A, B, K, P, expect)
}

func TestSample3(t *testing.T) {
	u, r := 30, 64
	A := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 1, 2}
	B := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 1, 2}
	K := []int{10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000}
	P := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 1}
	var expect int64 = 4758530000
	runSample(t, u, r, A, B, K, P, expect)
}
