package main

import (
	"math/rand"
	"sort"
	"testing"
)

func runSample(t *testing.T, x int, l int, r int, A []int) {
	res := solve(l, r, A)
	n := r - l + 1
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		y := A[i] ^ res
		arr[i] = y
	}

	sort.Ints(arr)

	for i := 0; i < n; i++ {
		if arr[i] != i {
			t.Fatalf("Sample %d, %v result %d, giving wrong result %v", x, A, res, arr)
		}
	}
}

func TestSample1(t *testing.T) {
	l, r := 0, 3
	A := []int{3, 2, 1, 0}
	runSample(t, 0, l, r, A)
}

func TestSample2(t *testing.T) {
	l, r := 0, 3
	A := []int{4, 7, 6, 5}
	runSample(t, 4, l, r, A)
}

func TestSample3(t *testing.T) {
	l, r := 0, 2
	A := []int{1, 2, 3}
	runSample(t, 3, l, r, A)
}

func TestSample4(t *testing.T) {
	n := 100
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	for j := 0; j < 10; j++ {
		x := rand.Intn(1 << (H - 1))
		for i := 0; i < n; i++ {
			arr[i] ^= x
		}
		runSample(t, x, 0, n-1, arr)
		for i := 0; i < n; i++ {
			arr[i] ^= x
		}
	}
}

func TestSample5(t *testing.T) {
	l, r := 0, 9
	A := []int{35225, 35224, 35227, 35226, 35229, 35228, 35231, 35230, 35217, 35216}
	runSample(t, 35225, l, r, A)
}

func TestSample6(t *testing.T) {
	l, r := 0, 9
	A := []int{43540, 43541, 43542, 43543, 43536, 43537, 43538, 43539, 43548, 43549}
	runSample(t, 43540, l, r, A)
}
