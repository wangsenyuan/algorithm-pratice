package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, P []int, s string, expect []int) {
	res := solve(P, s)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int{1, 2}
	s := "10"
	expect := []int{2, 1}
	runSample(t, P, s, expect)
}

func TestSample2(t *testing.T) {
	P := []int{3, 1, 2}
	s := "111"
	expect := []int{3, 1, 2}
	runSample(t, P, s, expect)
}

func TestSample3(t *testing.T) {
	P := []int{2, 3, 1, 8, 5, 4, 7, 6}
	s := "01110001"
	expect := []int{1, 6, 5, 8, 3, 2, 4, 7}
	runSample(t, P, s, expect)
}
