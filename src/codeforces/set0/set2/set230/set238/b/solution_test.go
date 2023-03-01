package main

import "testing"

func runSample(t *testing.T, h int, A []int, expect int) {
	res, arr := solve(A, h)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}

	var a int
	var b = 1 << 30

	n := len(A)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			tmp := A[i] + A[j]
			if arr[i] != arr[j] {
				tmp += h
			}
			if tmp > a {
				a = tmp
			}
			if tmp < b {
				b = tmp
			}
		}
	}
	if a-b != expect {
		t.Errorf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3}
	h := 2
	expect := 1
	runSample(t, h, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{0, 1, 0, 2, 1}
	h := 10
	expect := 3
	runSample(t, h, A, expect)
}
