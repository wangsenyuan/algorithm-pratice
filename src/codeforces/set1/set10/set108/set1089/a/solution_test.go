package main

import "testing"

func runSample(t *testing.T, a int, b int, expect int) {
	score, sets := solve(a, b)

	if expect == -inf {
		if len(score) > 0 {
			t.Fatalf("Sample expect %d, but got %v, %v", expect, score, sets)
		}
		return
	}

	if len(score) == 0 || score[0]-score[1] != expect {
		t.Fatalf("Sample expect %d, but got %v, %v", expect, score, sets)
	}

	check := func(cur []int, full int) bool {
		x, y := max(cur[0], cur[1]), min(cur[0], cur[1])
		if x == full {
			return y <= full-2
		}
		if x < full {
			return false
		}

		return x-y == 2
	}

	var x, y int
	var ans int

	for i, set := range sets {
		x += set[0]
		y += set[1]
		if i < 4 && !check(set, 25) {
			t.Fatalf("Sample result %v not correct", sets)
		}
		if i == 4 && !check(set, 15) {
			t.Fatalf("Sample result %v not correct", sets)
		}
		if set[0] > set[1] {
			ans++
		} else {
			ans--
		}
	}

	if x != a || y != b {
		t.Fatalf("Sample result %v not correct", sets)
	}
}

func TestSample1(t *testing.T) {
	a, b := 75, 0
	expect := 3
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a, b := 90, 90
	expect := 2
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a, b := 20, 0
	expect := -inf
	runSample(t, a, b, expect)
}

func TestSample4(t *testing.T) {
	a, b := 0, 75
	expect := -3
	runSample(t, a, b, expect)
}

func TestSample5(t *testing.T) {
	a, b := 78, 50
	expect := 3
	runSample(t, a, b, expect)
}

func TestSample6(t *testing.T) {
	a, b := 80, 100
	expect := 1
	runSample(t, a, b, expect)
}

func TestSample7(t *testing.T) {
	a, b := 0, 0
	expect := -inf
	runSample(t, a, b, expect)
}

func TestSample8(t *testing.T) {
	a, b := 25, 50
	expect := -inf
	runSample(t, a, b, expect)
}

func TestSample9(t *testing.T) {
	a, b := 27, 75
	expect := -2
	runSample(t, a, b, expect)
}
