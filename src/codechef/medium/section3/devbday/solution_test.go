package main

import "testing"

func runSample(t *testing.T, n int, gift []int, friends []int, expect int64) {
	res := solve(n, gift, friends)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, gift, friends, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	friends := []int{2, 1}
	gift := []int{3, 5}
	runSample(t, n, gift, friends, 8)
}

func TestSample2(t *testing.T) {
	n := 2
	friends := []int{2, 1}
	gift := []int{-3, 5}
	runSample(t, n, gift, friends, 2)
}

func TestSample3(t *testing.T) {
	n := 3
	friends := []int{2, 3, 2}
	gift := []int{-2, 2, 4}
	runSample(t, n, gift, friends, 6)
}

func TestSample4(t *testing.T) {
	n := 3
	friends := []int{2, 3, 2}
	gift := []int{-1, -2, -4}
	runSample(t, n, gift, friends, 0)
}

func TestSample5(t *testing.T) {
	n := 4
	friends := []int{2, 3, 2, 3}
	gift := []int{-2, 2, 4, 1}
	runSample(t, n, gift, friends, 7)
}

func TestSample6(t *testing.T) {
	n := 5
	friends := []int{5, 1, 4, 5, 4}
	gift := []int{-964684432, 74513386, -390331583, 309950241, 202556218}
	expect := int64(512506459)
	runSample(t, n, gift, friends, expect)
}
