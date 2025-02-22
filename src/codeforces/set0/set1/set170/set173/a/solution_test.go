package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A string, B string, expect []int) {
	res := solve(n, A, B)

	if !reflect.DeepEqual(expect, res) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	A := "RPS"
	B := "RSPP"
	expect := []int{3, 2}
	runSample(t, n, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := "RRRRRRRR"
	B := "R"
	expect := []int{0, 0}
	runSample(t, n, A, B, expect)
}
