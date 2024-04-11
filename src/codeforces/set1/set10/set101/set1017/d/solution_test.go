package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, w []int, s []string, x []string, k []int, expect []int) {
	res := solve(w, s, x, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	w := []int{40, 20}
	s := []string{
		"01",
		"01",
		"10",
		"11",
	}
	x := []string{
		"00",
		"00",
		"11",
		"11",
		"11",
	}
	k := []int{20, 40, 20, 40, 60}

	expect := []int{2, 4, 2, 3, 4}

	runSample(t, w, s, x, k, expect)
}

func TestSample2(t *testing.T) {
	w := []int{100}
	s := []string{
		"0", "1",
	}
	x := []string{
		"0", "0", "1", "1",
	}
	k := []int{0, 100, 0, 100}

	expect := []int{1, 2, 1, 2}

	runSample(t, w, s, x, k, expect)
}
