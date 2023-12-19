package main

import "testing"

func runSample(t *testing.T, l []int, expect int) {
	res := solve(l)

	if expect < 0 {
		if res >= 0 {
			t.Fatalf("Sample expect %d, but got %d", expect, res)
		}
		return
	}

	var liar int
	for i := 0; i < len(l); i++ {
		if l[i] > res {
			liar++
		}
	}
	if liar != res {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	l := []int{1, 2}
	expect := 1
	runSample(t, l, expect)
}

func TestSample2(t *testing.T) {
	l := []int{0, 0}
	expect := 0
	runSample(t, l, expect)
}

func TestSample3(t *testing.T) {
	l := []int{2, 2}
	expect := -1
	runSample(t, l, expect)
}

func TestSample4(t *testing.T) {
	l := []int{1}
	expect := -1
	runSample(t, l, expect)
}

func TestSample5(t *testing.T) {
	l := []int{0}
	expect := 0
	runSample(t, l, expect)
}

func TestSample6(t *testing.T) {
	l := []int{5, 5, 3, 3, 5}
	expect := 3
	runSample(t, l, expect)
}

func TestSample7(t *testing.T) {
	l := []int{5, 3, 6, 6, 3, 5}
	expect := 4
	runSample(t, l, expect)
}

func TestSample8(t *testing.T) {
	l := []int{0, 4, 1, 1, 7, 9, 8, 2, 2}
	expect := -1
	runSample(t, l, expect)
}
