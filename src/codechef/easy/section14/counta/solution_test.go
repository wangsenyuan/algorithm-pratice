package main

import "testing"

func runSample(t *testing.T, B []int, expect int) {
	res := solve(B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	B := []int{100}
	expect := 199801
	runSample(t, B, expect)
}

func TestSample2(t *testing.T) {
	B := []int{3, 9, 8, 4}
	expect := 199983
	runSample(t, B, expect)
}

func TestSample3(t *testing.T) {
	B := []int{10, 12}
	expect := 199977
	runSample(t, B, expect)
}

func TestSample4(t *testing.T) {
	B := []int{2, 5, 1, 1}
	expect := 999099867
	runSample(t, B, expect)
}