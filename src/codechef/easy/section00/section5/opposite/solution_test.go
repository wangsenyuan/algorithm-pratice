package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect bool, B []int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %t, but got %t", n, A, expect, res)
	} else if expect && !reflect.DeepEqual(A, B) {
		t.Errorf("Sample expect %v, but got %v", B, A)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{1, 1, 1, 1}
	B := []int{1, 1, 1, 1}
	runSample(t, n, A, true, B)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{1, 1, 1, 2}
	runSample(t, n, A, false, nil)
}

func TestSample3(t *testing.T) {
	n := 4
	A := []int{1, -1, -1, 4}
	B := []int{1, 4, 1, 4}
	runSample(t, n, A, true, B)
}

func TestSample4(t *testing.T) {
	n := 4
	A := []int{1, -1, 2, -1}
	runSample(t, n, A, false, nil)
}
