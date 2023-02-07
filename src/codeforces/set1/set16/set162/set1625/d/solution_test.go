package main

import "testing"

func runSample(t *testing.T, A []int, K int, expect int) {
	res := solve(A, K)

	if expect < 0 {
		if len(res) == 0 {
			return
		}
		t.Fatalf("Sample result %v, not correct", res)
	}

	if len(res) != expect {
		t.Fatalf("Sample result %v, but expect %d", res, expect)
	}

	for i := 0; i < expect; i++ {
		for j := i + 1; j < expect; j++ {
			if A[res[i]-1]^A[res[j]-1] < K {
				t.Fatalf("Sample result %v, A[%d] ^ A[%d] < %d not correct", res, res[i], res[j], K)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 8, 4, 16, 10, 14}
	K := 8
	expect := 3
	runSample(t, A, K, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 1, 4, 0}
	K := 1024
	expect := -1
	runSample(t, A, K, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6}
	K := 0
	expect := 6
	runSample(t, A, K, expect)
}

func TestSample5(t *testing.T) {
	A := []int{536872592, 1974}
	K := 908545581
	expect := -1
	runSample(t, A, K, expect)
}
