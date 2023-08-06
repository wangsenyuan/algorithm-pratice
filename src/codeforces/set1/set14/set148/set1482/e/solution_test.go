package main

import (
	rand2 "math/rand"
	"testing"
)

func runSample(t *testing.T, H []int, V []int) {
	expect := solve(H, V)
	res := solve1(H, V)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d for %v, %v", expect, res, H, V)
	}
}

func TestSample1(t *testing.T) {
	H := []int{1, 2, 3, 5, 4}
	V := []int{1, 5, 3, 2, 4}
	runSample(t, H, V)
}

func TestSample2(t *testing.T) {
	H := []int{1, 4, 3, 2, 5}
	V := []int{-3, 4, -10, 2, 7}
	runSample(t, H, V)
}

func TestSample3(t *testing.T) {
	H := []int{2, 1}
	V := []int{-2, -3}
	runSample(t, H, V)
}

func TestSample4(t *testing.T) {
	H := []int{4, 7, 3, 2, 5, 1, 9, 10, 6, 8}
	V := []int{-4, 40, -46, -8, -16, 4, -10, 41, 12, 3}
	runSample(t, H, V)
}

func TestSample5(t *testing.T) {
	V := []int{-4, 40, -46, -8, -16, 4, -10, 41, 12, 3}
	H := make([]int, len(V))
	for i := 0; i < len(H); i++ {
		H[i] = i + 1
	}
	for j := 0; j < 10; j++ {
		rand2.Shuffle(len(H), func(a, b int) {
			H[a], H[b] = H[b], H[a]
		})
		runSample(t, H, V)
	}
}

func TestSample6(t *testing.T) {
	V := []int{-4, 40, -46, -8, -16, 4, -10, 41, 12, 3}
	H := []int{7, 2, 9, 3, 8, 10, 5, 1, 6, 4}
	runSample(t, H, V)

}
