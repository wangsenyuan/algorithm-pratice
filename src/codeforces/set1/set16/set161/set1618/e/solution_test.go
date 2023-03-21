package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, B []int, expect []int) {
	res := solve(B)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	B := []int{12, 16, 14}
	expect := []int{3, 1, 3}
	runSample(t, B, expect)
}

func TestSample2(t *testing.T) {
	B := []int{1}
	expect := []int{1}
	runSample(t, B, expect)
}

func TestSample3(t *testing.T) {
	B := []int{1, 2, 3}
	//expect := []int{1}
	runSample(t, B, nil)
}

func TestSample4(t *testing.T) {
	B := []int{81, 75, 75, 93, 93, 87}
	expect := []int{5, 5, 4, 1, 4, 5}
	runSample(t, B, expect)
}
