package main

import (
	"math/rand"
	"testing"
)

func runSample(t *testing.T, H []int, expect int) {
	res := solve(H)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d for %v ", expect, res, H)
	}
}

func TestSample1(t *testing.T) {
	h := []int{1, 3, 1, 4, 5}
	expect := 3
	runSample(t, h, expect)
}

func TestSample2(t *testing.T) {
	h := []int{4, 2, 2, 4}
	expect := 1
	runSample(t, h, expect)
}

func TestSample3(t *testing.T) {
	h := []int{1, 2, 3, 4}
	expect := 3
	runSample(t, h, expect)
}

func TestSample5(t *testing.T) {
	h := []int{100, 1, 100, 1, 100}
	expect := 2
	runSample(t, h, expect)
}

func TestSample6(t *testing.T) {
	h := []int{4, 3, 2, 1}
	expect := 3
	runSample(t, h, expect)
}

func TestSample7(t *testing.T) {
	n := 11
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100)
	}
	expect := bruteForce(arr)
	runSample(t, arr, expect)
}

func TestSample8(t *testing.T) {
	arr := []int{61, 78, 64, 76, 97, 7, 97, 18, 7, 22, 85}
	expect := bruteForce(arr)
	runSample(t, arr, expect)
}

func TestSample10(t *testing.T) {
	n := 13
	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = 7
	}
	arr[n-1] = 6
	expect := bruteForce(arr)
	runSample(t, arr, expect)
}
