package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, H []int, W []int, X []string, M int, Q [][]int, expect []string) {
	res := solve(N, H, W, X, M, Q)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v %v %d %v, expect %v, but got %v", N, H, W, X, M, Q, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 5
	H := []int{1000, 2000, 2000, 1000, 1500}
	W := []int{1000, 1000, 2000, 2000, 2010}
	X := []string{"BIRD", "BIRD", "BIRD", "BIRD", "NOT BIRD"}

	M := 3
	Q := [][]int{
		{1500, 1500},
		{900, 900},
		{1400, 2020},
	}
	expect := []string{"BIRD", "UNKNOWN", "NOT BIRD"}
	runSample(t, N, H, W, X, M, Q, expect)
}

func TestSample2(t *testing.T) {
	N := 3
	H := []int{500, 501, 502}
	W := []int{700, 700, 700}
	X := []string{"NOT BIRD", "BIRD", "NOT BIRD"}

	M := 2
	Q := [][]int{
		{501, 600},
		{502, 501},
	}
	expect := []string{"UNKNOWN", "NOT BIRD"}
	runSample(t, N, H, W, X, M, Q, expect)
}

func TestSample3(t *testing.T) {
	N := 1
	H := []int{100}
	W := []int{100}
	X := []string{"NOT BIRD"}

	M := 3
	Q := [][]int{
		{107, 93},
		{86, 70},
		{110, 115},
	}
	expect := []string{"UNKNOWN", "UNKNOWN", "UNKNOWN"}
	runSample(t, N, H, W, X, M, Q, expect)
}
