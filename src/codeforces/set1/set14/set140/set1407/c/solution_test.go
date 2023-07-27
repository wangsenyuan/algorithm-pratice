package main

import (
	rand2 "math/rand"
	"reflect"
	"testing"
)

func runSample(t *testing.T, arr []int) {
	var cnt int
	ask := func(i int, j int) int {
		cnt++
		return arr[i-1] % arr[j-1]
	}

	res := solve(len(arr), ask)

	if cnt > 2*len(arr) {
		t.Fatalf("Sample asked too much %d", cnt)
	}

	if !reflect.DeepEqual(res, arr) {
		t.Fatalf("Sample expect %v, but got %v", arr, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 3, 2}
	runSample(t, arr)
}

func TestSample2(t *testing.T) {
	n := 10000
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	rand2.Shuffle(n, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	runSample(t, arr)
}
