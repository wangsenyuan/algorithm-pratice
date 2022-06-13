package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, L int, R int, expect []int) {
	res := solve(N, L, R)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, L, R := 2, -2, 2
	expect := []int{1, 0, 2, 0, 1}
	runSample(t, N, L, R, expect)
}

func TestSample3(t *testing.T) {
	N, L, R := 2, 0, 2
	expect := []int{2, 0, 1}
	runSample(t, N, L, R, expect)
}

func TestSample2(t *testing.T) {
	N, L, R := 150000, 48, 48
	expect := []int{122830846}
	runSample(t, N, L, R, expect)
}

func TestSample4(t *testing.T) {
	N, L, R := 0, 0, 0
	expect := []int{1}
	runSample(t, N, L, R, expect)
}
