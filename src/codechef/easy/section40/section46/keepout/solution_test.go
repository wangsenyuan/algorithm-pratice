package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, m int, A []int, expect []int) {
	res := solve(m, A)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 5
	A := []int{2, 4}
	expect := []int{3, 2}
	runSample(t, m, A, expect)
}

func TestSample2(t *testing.T) {
	m := 5
	A := []int{4, 3, 2, 1}
	expect := []int{4, 3, 2, 1}
	runSample(t, m, A, expect)
}

func TestSample3(t *testing.T) {
	m := 25
	A := []int{10, 12, 5, 19}
	expect := []int{15, 13, 13, 7}
	runSample(t, m, A, expect)
}
