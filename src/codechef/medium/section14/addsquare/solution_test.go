package main

import "testing"

func TestBitSet(t *testing.T) {
	n := 64

	a := NewBitSet(n)

	a.Set(10)

	if !a.IsSet(10) {
		t.Fatalf("10 should be set, but not")
	}
	a.Clear(10)

	if a.IsSet(10) {
		t.Fatalf("10 should not be set but wrong")
	}
	a.Set(10)
	a = a.RightShift(3)

	if a.IsSet(10) || !a.IsSet(7) {
		t.Fatalf("7 should be set")
	}
	a = a.RightShift(8)

	if a.IsSet(0) {
		t.Fatalf("no one should be set")
	}
}

func TestSample1(t *testing.T) {
	W, H, N, M := 10, 10, 3, 3
	X := []int{3, 6, 8}
	Y := []int{1, 6, 10}
	res := solve(W, H, N, M, X, Y)

	if res != 3 {
		t.Errorf("Sample expect 3, but got %d", res)
	}
}
