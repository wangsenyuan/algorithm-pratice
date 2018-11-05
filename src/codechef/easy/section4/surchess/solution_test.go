package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, m int, A []string, Q int, quries []int, expect []int) {
	res := solve(n, m, A, Q, quries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %d %v, expect %v, but got %v", n, m, A, Q, quries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 8, 8
	A := []string{
		"00101010",
		"00010101",
		"10101010",
		"01010101",
		"10101010",
		"01010101",
		"10101010",
		"01010101",
	}
	Q := 4
	quries := []int{1, 2, 0, 1001}

	expect := []int{7, 8, 6, 8}
	runSample(t, n, m, A, Q, quries, expect)
}
