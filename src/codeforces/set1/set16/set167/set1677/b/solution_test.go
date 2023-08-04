package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, s string, expect []int) {
	res := solve(n, m, s)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 2
	s := "1100"
	expect := []int{2, 3, 4, 3}
	runSample(t, n, m, s, expect)
}

func TestSample2(t *testing.T) {
	n, m := 4, 2
	s := "11001101"
	expect := []int{2, 3, 4, 3, 5, 4, 6, 5}
	runSample(t, n, m, s, expect)
}

func TestSample3(t *testing.T) {
	n, m := 2, 4
	s := "11001101"
	expect := []int{2, 3, 3, 3, 4, 4, 4, 5}
	runSample(t, n, m, s, expect)
}
