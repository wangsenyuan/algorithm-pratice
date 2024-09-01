package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, expect []int) {
	res := solve(n)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	expect := []int{1}
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	expect := []int{1, 2}
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	expect := []int{2, 1, 3}
	runSample(t, n, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	expect := []int{3, 1, 2, 4}
	runSample(t, n, expect)
}

func TestSample5(t *testing.T) {
	n := 5
	expect := []int{2, 4, 1, 3, 5}
	runSample(t, n, expect)
}

func TestSample6(t *testing.T) {
	n := 6
	expect := []int{3, 4, 1, 5, 2, 6}
	runSample(t, n, expect)
}
