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
	p := []int{2, 1, 4, 3, 6, 5}
	expect := []int{1, 3}
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := []int{2, 3, 1, 5, 6, 4}
	expect := []int{2, 2}
	runSample(t, p, expect)
}

func TestSample3(t *testing.T) {
	p := []int{2, 3, 2, 5, 6, 5, 8, 9, 8}
	expect := []int{1, 3}
	runSample(t, p, expect)
}

func TestSample4(t *testing.T) {
	p := []int{2, 1}
	expect := []int{1, 1}
	runSample(t, p, expect)
}

func TestSample5(t *testing.T) {
	p := []int{4, 3, 2, 1}
	expect := []int{1, 2}
	runSample(t, p, expect)
}

func TestSample6(t *testing.T) {
	p := []int{2, 3, 4, 5, 1}
	expect := []int{1, 1}
	runSample(t, p, expect)
}

func TestSample7(t *testing.T) {
	p := []int{5, 3, 4, 1, 1, 2}
	expect := []int{1, 1}
	runSample(t, p, expect)
}

func TestSample8(t *testing.T) {
	p := []int{3, 5, 4, 1, 2}
	expect := []int{2, 2}
	runSample(t, p, expect)
}

func TestSample9(t *testing.T) {
	p := []int{6, 3, 2, 5, 4, 3}
	expect := []int{1, 2}
	runSample(t, p, expect)
}

func TestSample10(t *testing.T) {
	p := []int{5, 1, 4, 3, 4, 2}
	expect := []int{1, 1}
	runSample(t, p, expect)
}
