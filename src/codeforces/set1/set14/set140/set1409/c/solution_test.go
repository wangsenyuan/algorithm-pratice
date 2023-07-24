package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, arr []int, x int, y int) {
	res := solve(len(arr), x, y)
	sort.Ints(arr)
	if !reflect.DeepEqual(res, arr) {
		t.Fatalf("Sample expect %v, but got %v", arr, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 49}
	x, y := 1, 49
	runSample(t, arr, x, y)
}

func TestSample2(t *testing.T) {
	arr := []int{20, 40, 30, 50, 10}
	x, y := 20, 50
	runSample(t, arr, x, y)
}

func TestSample3(t *testing.T) {
	arr := []int{26, 32, 20, 38, 44, 50}
	x, y := 20, 50
	runSample(t, arr, x, y)
}

func TestSample4(t *testing.T) {
	arr := []int{8, 23, 18, 13, 3}
	x, y := 3, 8
	runSample(t, arr, x, y)
}
