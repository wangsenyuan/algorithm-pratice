package main

import "testing"

func runSample(t *testing.T, h []int, add int, rem int, move int, expect int) {
	res := solve(h, add, rem, move)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	// 6, 6, 6
	// 5, 4, 8
	h := []int{1, 3, 8}
	add, rem, move := 1, 100, 100
	expect := 12
	runSample(t, h, add, rem, move, expect)
}

func TestSample2(t *testing.T) {
	// 6, 6, 6
	// 5, 4, 8
	h := []int{1, 3, 8}
	add, rem, move := 100, 1, 100
	expect := 9
	runSample(t, h, add, rem, move, expect)
}

func TestSample3(t *testing.T) {
	// 6, 6, 6
	// 5, 4, 8
	h := []int{1, 3, 8}
	add, rem, move := 100, 100, 1
	expect := 4
	runSample(t, h, add, rem, move, expect)
}

func TestSample4(t *testing.T) {
	// 6, 6, 6
	// 5, 4, 8
	h := []int{5, 5, 3, 6, 5}
	add, rem, move := 1, 2, 4
	expect := 4
	runSample(t, h, add, rem, move, expect)
}

func TestSample5(t *testing.T) {
	// 6, 6, 6
	// 5, 4, 8
	h := []int{5, 5, 3, 6, 5}
	add, rem, move := 1, 2, 2
	expect := 3
	runSample(t, h, add, rem, move, expect)
}

func TestSample6(t *testing.T) {
	// 6, 6, 6
	// 5, 4, 8
	h := []int{3, 10, 4, 9, 2, 7, 6, 10, 4, 8}
	add, rem, move := 7, 8, 3
	expect := 57
	runSample(t, h, add, rem, move, expect)
}
