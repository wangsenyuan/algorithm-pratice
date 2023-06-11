package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, Q []int, expect [][]int) {
	res := solve(A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{95, 143, 58, 119, 69}
	Q := []int{1, 3, 5}
	expect := [][]int{
		{469693188, 309321521},
		{298495478, 683994827},
		{620264520, 762751641},
	}
	runSample(t, A, Q, expect)
}
