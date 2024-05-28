package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1}
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1}
	expect := 2
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 1}
	expect := 1
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{2, 2}
	expect := 1
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{1, 1, 1}
	expect := 3
	runSample(t, a, expect)
}
func TestSample6(t *testing.T) {
	a := []int{3, 1, 2}
	expect := 1
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{1, 3, 1}
	expect := 2
	runSample(t, a, expect)
}

func TestSample8(t *testing.T) {
	a := []int{1, 2, 3, 1, 2, 4, 2}
	expect := 3
	runSample(t, a, expect)
}

func TestSample9(t *testing.T) {
	a := []int{2, 1, 1, 1, 2, 3, 1}
	expect := 4
	runSample(t, a, expect)
}

func TestSample10(t *testing.T) {
	a := []int{2, 2, 5, 1, 6, 1, 8, 2, 8, 2}
	// 这个确实是2， 在5处向右，在7处向左
	expect := 2
	runSample(t, a, expect)
}

func TestSample11(t *testing.T) {
	a := []int{2, 1, 2, 1, 1, 2}
	expect := 3
	runSample(t, a, expect)
}

func TestSample12(t *testing.T) {
	a := []int{1, 1, 4, 1, 3, 2}
	expect := 3
	runSample(t, a, expect)
}

func TestSample13(t *testing.T) {
	a := []int{1, 1, 1, 4, 1, 1, 1}
	expect := 4
	runSample(t, a, expect)
}
