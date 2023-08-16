package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, k int, expect []int) {
	res := solve(A, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	A := []int{11, 11, 12, 13, 13, 14, 14}
	expect := []int{13, 14}
	runSample(t, A, k, expect)
}

func TestSample2(t *testing.T) {
	k := 1
	A := []int{6, 3, 5, 2, 1}
	expect := []int{1, 3}
	runSample(t, A, k, expect)
}

func TestSample3(t *testing.T) {
	k := 4
	A := []int{4, 3, 4, 3, 3, 4}
	//expect := []int{1, 3}
	runSample(t, A, k, nil)
}

func TestSample4(t *testing.T) {
	k := 2
	A := []int{1, 1, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4}
	expect := []int{1, 4}
	runSample(t, A, k, expect)
}
