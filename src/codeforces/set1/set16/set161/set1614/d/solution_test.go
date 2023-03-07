package main

import "testing"

func runSample(t *testing.T, A []int, expect int64) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 3, 1, 2, 6, 2}
	var expect int64 = 14
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 1}
	var expect int64 = 3
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 2, 2}
	var expect int64 = 6
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{5, 7, 10, 3, 1, 10, 100, 3, 42, 54}
	var expect int64 = 131
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{133, 10374, 4389, 132, 228, 16302, 6, 627, 924, 20748, 399, 14, 2717, 10868, 2926, 28, 182, 1729, 858,
		21, 19019, 8151, 2, 38038, 494, 247, 3458, 6916, 3003, 266, 5187, 32604, 5434, 42, 532, 12, 76, 1596, 4004,
		364, 143, 114114, 114, 22, 273, 1463, 38, 78, 4, 836, 209, 418, 1716, 57, 231, 2002, 26, 2964, 1001, 156, 11,
		13, 8778, 308, 76076, 572, 1482, 1092, 798, 7, 2508, 6006, 57057, 91, 77, 39, 33, 741, 988, 286, 19, 546, 12012,
		154, 84, 228228, 3, 17556, 5852, 429, 1254, 66, 44, 462, 1, 52,
	}
	var expect int64 = 476226
	runSample(t, A, expect)
}

func TestSample6(t *testing.T) {
	A := []int{21, 24, 27}
	var expect int64 = 33
	runSample(t, A, expect)
}
