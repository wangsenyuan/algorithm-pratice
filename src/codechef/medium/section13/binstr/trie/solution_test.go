package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, Q int, A []string, L []int, R []int, X []string, expect []int) {
	res := solve(N, Q, A, L, R, X)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, Q := 5, 4
	A := []string{
		"100",
		"101",
		"1",
		"1011",
		"11",
	}
	L := []int{2, 1, 3, 1}
	R := []int{3, 5, 5, 5}
	X := []string{"10", "1100", "1010", "11100"}
	expect := []int{2, 5, 3, 5}
	runSample(t, N, Q, A, L, R, X, expect)
}
