package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res, steps := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
	if expect < 0 {
		return
	}

	if len(steps) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, steps)
	}
	pos := len(a)
	for _, j := range steps {
		if j < pos-a[pos-1] {
			t.Fatalf("cant jump %d from %d", j, pos)
		}
		// j >= pos - a[pos-1]
		if j == 0 {
			break
		}
		pos = j + b[j-1]
	}
}

func TestSample1(t *testing.T) {
	a := []int{0, 2, 2}
	b := []int{1, 1, 0}
	expect := 2
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1}
	b := []int{1, 0}
	expect := -1
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{0, 1, 2, 3, 5, 5, 6, 7, 8, 5}
	b := []int{9, 8, 7, 1, 5, 4, 3, 2, 0, 0}
	expect := 3
	runSample(t, a, b, expect)
}

func TestSample4(t *testing.T) {
	a := []int{0, 0, 1, 4, 2}
	b := []int{0, 2, 0, 1, 0}
	expect := 3
	runSample(t, a, b, expect)
}
