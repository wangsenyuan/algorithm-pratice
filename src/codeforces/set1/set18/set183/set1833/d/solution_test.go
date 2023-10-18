package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, p []int, expect []int) {
	res := solve(p)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{2, 3, 1, 5, 4}
	expect := []int{5, 4, 1, 3, 2}
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := []int{4, 1, 6, 7, 2, 8, 5, 3, 9}
	expect := []int{9, 4, 1, 6, 7, 2, 8, 5, 3}
	runSample(t, p, expect)
}

func TestSample3(t *testing.T) {
	p := []int{4, 3, 2, 1}
	expect := []int{3, 2, 1, 4}
	runSample(t, p, expect)
}

func TestSample4(t *testing.T) {
	p := []int{2, 1}
	expect := []int{1, 2}
	runSample(t, p, expect)
}

func TestSample5(t *testing.T) {
	p := []int{3, 2, 4, 1, 5, 6}
	expect := []int{6, 5, 3, 2, 4, 1}
	runSample(t, p, expect)
}

func TestSample6(t *testing.T) {
	p := []int{3, 2, 1, 5, 7, 6, 4}
	expect := []int{7, 6, 4, 5, 3, 2, 1}
	runSample(t, p, expect)
}

func TestSample7(t *testing.T) {
	p := []int{10, 2, 5, 6, 1, 9, 3, 8, 4, 7}
	expect := []int{9, 3, 8, 4, 7, 1, 10, 2, 5, 6}
	runSample(t, p, expect)
}

func TestSample8(t *testing.T) {
	p := []int{4, 2, 1, 3}
	expect := []int{3, 4, 2, 1}
	runSample(t, p, expect)
}

func TestSample9(t *testing.T) {
	p := []int{1}
	expect := []int{1}
	runSample(t, p, expect)
}
