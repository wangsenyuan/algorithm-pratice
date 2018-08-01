package main

import "testing"

func runSample(t *testing.T, B []int, expect int) {
	res := solve(B)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	B := []int{2, 3}
	runSample(t, B, 3)
}

func TestSample2(t *testing.T) {
	B := []int{2, 3, 7}
	runSample(t, B, 143)
}

func TestSample3(t *testing.T) {
	B := []int{9, 10}
	runSample(t, B, 91)
}

func TestSample4(t *testing.T) {
	B := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	runSample(t, B, 11814485)
}

func TestCheck(t *testing.T) {
	res := happy(143, 7)
	if !res {
		t.Errorf("check 143, 7, should got true, but got false")
	}
	res = happy(143, 3)
	if !res {
		t.Errorf("check 143, 3, should got true, but got false")
	}

	res = happy(91, 9)
	if !res {
		t.Errorf("check 91, 9, should got true, but got false")
	}

	res = happy(91, 10)
	if !res {
		t.Errorf("check 91, 10, should got true, but got false")
	}
}
