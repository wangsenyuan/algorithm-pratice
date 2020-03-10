package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, Q int, queries [][]int, expect []int64) {
	res := solve(n, A, Q, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v, expect %v, but got %v", n, A, Q, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, Q := 3, 2
	A := []int{1, 2, 3}
	queries := [][]int{
		{2, 2},
		{1, 3},
	}
	expect := []int64{0, 3}
	runSample(t, n, A, Q, queries, expect)
}

func TestBitRange(t *testing.T) {
	// supporse we want to do range update/query on arr [0, 1, 2, 3, 4]
	bit := NewBIT2(5)
	bit.UpdateRange(5, 0, 4, 1)
	sum := bit.SumRange(5, 4)
	if sum != 5 {
		t.Errorf("Sample expect 5, but got %d", sum)
	}
	bit.UpdateRange(5, 1, 4, 1)

	sum = bit.SumRange(5, 4)
	if sum != 9 {
		t.Errorf("Sample expect 9, but got %d", sum)
	}

	bit.UpdateRange(5, 2, 4, 1)

	sum = bit.SumRange(5, 4)
	if sum != 12 {
		t.Errorf("Sample expect 12, but got %d", sum)
	}

	bit.UpdateRange(5, 3, 4, 1)

	sum = bit.SumRange(5, 4)
	if sum != 14 {
		t.Errorf("Sample expect 14, but got %d", sum)
	}

	bit.UpdateRange(5, 4, 4, 1)
	sum = bit.SumRange(5, 4)
	if sum != 15 {
		t.Errorf("Sample expect 15, but got %d", sum)
	}
	bit.UpdateRange(5, 1, 1, 5)
	sum = bit.SumRange(5, 4)
	if sum != 20 {
		t.Errorf("Sample expect 20, but got %d", sum)
	}
}
