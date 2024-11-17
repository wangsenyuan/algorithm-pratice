package p3357

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := minDifference(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, -1, 10, 8}
	expect := 4
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{-1, -1, -1}
	expect := 0
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{-1, 10, -1, 8}
	expect := 1
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 12}
	expect := 11
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{-1, -1, 13, 34}
	expect := 21
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{20, -1, 72, -1, 108}
	expect := 26
	runSample(t, a, expect)
}
