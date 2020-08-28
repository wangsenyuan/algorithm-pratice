package main

import (
	"math/rand"
	"testing"
	"time"
)

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{1, 2, 4}
	expect := 6
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{5, 5, 5, 5, 5}
	expect := 0
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 1000)
	compareSolution(t, arr)
}

func compareSolution(t *testing.T, arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(MOD)
	}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	expect := solve1(len(arr), arr)
	res := solve(len(arr), arr)
	if expect != res {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample4(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for n := 10; n < 10000; n *= 10 {
		arr := make([]int, n)
		compareSolution(t, arr)
	}
}
