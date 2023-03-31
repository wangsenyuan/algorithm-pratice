package main

import (
	"math/rand"
	"testing"
)

func runSample(t *testing.T, arr []int) {
	var cnt int
	ask := func(i, j, k int) int {
		a := max(arr[i-1], arr[j-1], arr[k-1])
		b := min(arr[i-1], arr[j-1], arr[k-1])
		cnt++
		return a - b
	}
	n := len(arr)

	res := solve(n, ask)

	if cnt > 2*len(arr)-2 {
		t.Fatalf("check took %d times > %d", cnt, 2*len(arr)-2)
	}

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			if res[0] == i+1 || res[1] == i+1 {
				//t.Logf("Sample result %v, correct", res)
				return
			}
			t.Fatalf("Sample result %v, not correct, as 0 at %d in %d", res, i, n)
		}
	}
}

func TestSample1(t *testing.T) {
	arr := []int{0, 1, 2, 3}
	runSample(t, arr)
}

func runSampleN(t *testing.T, n int) {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
	}

	rng := rand.New(rand.NewSource(17))

	for i := n - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	runSample(t, arr)
}

func TestSample2(t *testing.T) {
	runSampleN(t, 1000)
}

func TestSample3(t *testing.T) {
	arr := []int{2, 2, 2, 2, 0}
	runSample(t, arr)
}

func TestSample4(t *testing.T) {
	for n := 10; n < 1000; n += 11 {
		runSampleN(t, n)
	}
}

func TestSample5(t *testing.T) {
	runSampleN(t, 10)
}
