package main

import "testing"

func runSample(t *testing.T, A []int, k int, x, y int) {
	res, segs := solve(A, k)

	if y-x != res[1]-res[0] {
		t.Fatalf("Sample result %v, %v, not correct", res, segs)
	}
	var m int
	for _, seg := range segs {
		a, b := seg[0], seg[1]
		var cnt int
		for i := a - 1; i < b; i++ {
			if A[i] >= res[0] && A[i] <= res[1] {
				cnt++
			}
		}
		if cnt <= b-a+1-cnt {
			t.Fatalf("Sample result %v, not correct", segs)
		}
		m += b - a + 1
	}

	if m != len(A) {
		t.Fatalf("Sample result %v, not correct", segs)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	A := []int{1, 2}
	x, y := 1, 2
	runSample(t, A, k, x, y)
}

func TestSample2(t *testing.T) {
	k := 2
	A := []int{1, 2, 2, 2}
	x, y := 2, 2
	runSample(t, A, k, x, y)
}

func TestSample3(t *testing.T) {
	k := 3
	A := []int{5, 5, 5, 1, 5, 5, 1, 5, 5, 5, 1}
	x, y := 5, 5
	runSample(t, A, k, x, y)
}

func TestSample4(t *testing.T) {
	k := 1
	A := []int{1}
	x, y := 1, 1
	runSample(t, A, k, x, y)
}
