package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, K int, expect []int) {
	res := solve(A, K)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 1, 3}
	K := 1
	expect := []int{3, 2, 1}
	runSample(t, A, K, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 1, 3}
	K := 2
	expect := []int{1, 3, 2}
	runSample(t, A, K, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 2, 3}
	K := 1
	runSample(t, A, K, nil)
}

func TestSample4(t *testing.T) {
	A := []int{1, 2, 3}
	K := 3
	expect := []int{1, 2, 3}
	runSample(t, A, K, expect)
}
