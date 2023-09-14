package main

import "testing"

func runSample(t *testing.T, n int, crickets []int, x int, y int, expect bool) {
	res := solve(n, crickets, x, y)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 8
	crickets := []int{
		7, 2, 8, 2, 7, 1,
	}
	x, y := 5, 1
	expect := true
	runSample(t, n, crickets, x, y, expect)
}

func TestSample2(t *testing.T) {
	n := 8
	crickets := []int{
		2, 2, 1, 2, 2, 1,
	}
	x, y := 5, 5
	expect := false
	runSample(t, n, crickets, x, y, expect)
}

func TestSample3(t *testing.T) {
	n := 8
	crickets := []int{
		2, 2, 1, 2, 2, 1,
	}
	x, y := 6, 6
	expect := true
	runSample(t, n, crickets, x, y, expect)
}

func TestSample4(t *testing.T) {
	n := 8
	crickets := []int{
		1, 1, 1, 2, 2, 1,
	}
	x, y := 5, 5
	expect := false
	runSample(t, n, crickets, x, y, expect)
}

func TestSample5(t *testing.T) {
	n := 8
	crickets := []int{
		2, 2, 1, 2, 2, 1,
	}
	x, y := 8, 8
	expect := true
	runSample(t, n, crickets, x, y, expect)
}

func TestSample6(t *testing.T) {
	n := 8
	crickets := []int{
		8, 8, 8, 7, 7, 8,
	}
	x, y := 4, 8
	expect := true
	runSample(t, n, crickets, x, y, expect)
}
