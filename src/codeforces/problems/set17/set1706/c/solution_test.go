package main

import "testing"

func runSample(t *testing.T, H []int, expect int64) {
	res := solve(H)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	H := []int{2, 1, 2}
	expect := int64(2)
	runSample(t, H, expect)
}

func TestSample2(t *testing.T) {
	H := []int{1, 2, 1, 4, 3}
	expect := int64(0)
	runSample(t, H, expect)
}

func TestSample3(t *testing.T) {
	H := []int{3, 1, 4, 5, 5, 2}
	expect := int64(3)
	runSample(t, H, expect)
}

func TestSample4(t *testing.T) {
	H := []int{4, 2, 1, 3, 5, 3, 6, 1}
	expect := int64(3)
	runSample(t, H, expect)
}
func TestSample5(t *testing.T) {
	H := []int{1, 10, 1, 1, 10, 1}
	expect := int64(0)
	runSample(t, H, expect)
}
