package main

import "testing"

func runSample(t *testing.T, A []int, k int) {
	ask := func(l, r int) int {
		return sumRange(A, l, r)
	}
	l, r := solve(len(A), k, ask)
	sum := sumRange(A, l-1, r-1)

	if sum != k {
		t.Errorf("Sample result %d %d, not correct", l, r)
	}

}

func sumRange(arr []int, l int, r int) int {
	var sum int
	for i := l; i <= r; i++ {
		sum += arr[i]
	}
	return sum
}

func TestSample1(t *testing.T) {
	A := []int{1, 1, 2, 2}
	k := 5
	runSample(t, A, k)
}

func TestSample2(t *testing.T) {
	A := []int{2, 2, 1, 2}
	k := 5
	runSample(t, A, k)
}

func TestSample3(t *testing.T) {
	A := []int{2, 2, 2, 1, 1}
	k := 5
	runSample(t, A, k)
}

func TestSample4(t *testing.T) {
	A := []int{2, 2, 2, 1}
	k := 5
	runSample(t, A, k)
}
