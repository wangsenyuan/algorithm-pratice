package main

import "testing"

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func TestSample1(t *testing.T) {
	n := 1
	as := []int{3}
	bs := []int{4}
	res := solve(n, as, bs)
	if abs(res-1.0) > 0.000001 {
		t.Errorf("this sample should give answer 1.0, but got %.6f", res)
	}
}

func TestSample2(t *testing.T) {
	n := 2
	as := []int{2, 3}
	bs := []int{1, 4}
	res := solve(n, as, bs)
	if abs(res-1.5) > 0.000001 {
		t.Errorf("this sample should give answer 1.5, but got %.6f", res)
	}
}

func TestSample3(t *testing.T) {
	n := 2
	as := []int{2, 4}
	bs := []int{2, 2}
	res := solve(n, as, bs)
	if abs(res-0) > 0.000001 {
		t.Errorf("this sample should give answer 1.5, but got %.6f", res)
	}
}
