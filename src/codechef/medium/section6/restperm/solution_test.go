package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func runSample(t *testing.T, arr []int) {
	div := func(x, y, d int) bool {
		return abs(arr[x]-arr[y])%d == 0
	}
	less := func(x, y int) bool {
		return arr[x] < arr[y]
	}
	res := solve(len(arr), div, less)

	if !reflect.DeepEqual(res, arr) {
		t.Errorf("Sample expect %v, but got %v", arr, res)
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func TestSample1(t *testing.T) {
	arr := []int{1, 2, 3}
	runSample(t, arr)
}

func TestSample2(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	runSample(t, arr)
}

func TestSample3(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	runSample(t, arr)
}

func TestSample4(t *testing.T) {
	arr := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i + 1
	}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	runSample(t, arr)
}
