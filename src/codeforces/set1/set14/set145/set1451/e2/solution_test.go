package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func runSample(t *testing.T, arr []int) {
	var cnt int
	ask := func(tp string, i int, j int) int {
		cnt++
		switch tp {
		case "XOR":
			return arr[i-1] ^ arr[j-1]
		case "OR":
			return arr[i-1] | arr[j-1]
		default:
			return arr[i-1] & arr[j-1]
		}
	}

	res := solve(len(arr), ask)

	if cnt > len(res)+1 {
		t.Fatalf("Sample expect no more than %d queries, but got %d", len(res)+1, cnt)
	}

	if !reflect.DeepEqual(res, arr) {
		t.Fatalf("Sample expect %v, but got %v", arr, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{0, 0, 2, 3}
	runSample(t, arr)
}

func TestSample2(t *testing.T) {
	n := 65536
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	rand.Shuffle(n, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	runSample(t, arr)
}

func TestSample3(t *testing.T) {
	arr := []int{2, 7, 0, 15, 11, 1, 5, 8, 6, 4, 9, 12, 3, 10, 14, 13}

	runSample(t, arr)
}

func TestSample4(t *testing.T) {
	arr := []int{9, 6, 12, 15, 5, 14, 0, 2, 8, 7, 10, 11, 1, 4, 13, 3}

	runSample(t, arr)
}
