package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, blocks []int, expect []int) {
	res := solve(blocks)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	blocks := []int{1, 2, 3, 1, 2, 3, 1}
	expect := []int{3, 2, 2, 0, 0, 0, 0}
	runSample(t, blocks, expect)
}

func TestSample2(t *testing.T) {
	blocks := []int{4, 2, 2, 2, 4, 4}
	expect := []int{0, 3, 0, 2, 0, 0}
	runSample(t, blocks, expect)
}

func TestSample3(t *testing.T) {
	blocks := []int{1}
	expect := []int{1}
	runSample(t, blocks, expect)
}

func TestSample4(t *testing.T) {
	blocks := []int{5, 4, 5, 3, 5}
	expect := []int{0, 0, 1, 1, 1}
	runSample(t, blocks, expect)
}

func TestSample5(t *testing.T) {
	blocks := []int{3, 3, 3, 1, 3, 3}
	expect := []int{1, 0, 4, 0, 0, 0}
	runSample(t, blocks, expect)
}

func TestSample6(t *testing.T) {
	blocks := []int{1, 2, 3, 4, 4, 3, 2, 1}
	expect := []int{2, 2, 2, 2, 0, 0, 0, 0}
	runSample(t, blocks, expect)
}
