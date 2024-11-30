package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}

	// len(res) == expect
	var s int

	for j := 0; j < len(res); j++ {
		var sum int
		for i := res[j][0] - 1; i < res[j][1]; i++ {
			sum += a[i]
		}
		if j == 0 {
			s = sum
		} else if s != sum {
			t.Fatalf("Sample result %v, not correct, it has different sum at %v", res, res[j])
		}
	}
}

func TestSample1(t *testing.T) {
	a := []int{4, 1, 2, 2, 1, 5, 3}
	expect := 3
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
	expect := 2
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 1, 1, 1}
	expect := 4
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{-7, -8, -15, 13, -5, -7, -4, -5, -8, 29, -3, -2, 7, -12, 3, -2, -9, 4, 24, 1, -11, 14, -2, -13, -15, -3, 18, -1, 4}
	expect := 6
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{4, 1, 2, 1, 5, 3}
	expect := 2
	runSample(t, a, expect)
}
