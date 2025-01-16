package main

import "testing"

func runSample(t *testing.T, m int, x int, p []int) {
	var found bool
	var id int
	n := len(p)
	ask := func(y int) int {
		if y <= 0 || y > m {
			t.Fatalf("Sample asked %d, out ouf bounds", y)
		}
		id++
		if id > 60 {
			t.Fatalf("Sample asked too much")
		}
		if y == x {
			if found {
				t.Fatalf("already found the the answer, but don't stop")
			}
			found = true
			return 0
		}
		ans := 1
		if x < y {
			ans = -1
		}
		if p[(id-1)%n] == 0 {
			ans *= -1
		}
		return ans
	}

	solve(m, n, ask)

	if !found {
		t.Fatalf("Sample solution not found the correct answer %d", x)
	}
}

func TestSample1(t *testing.T) {
	m, x := 5, 3
	p := []int{1, 0}
	runSample(t, m, x, p)
}

func TestSample2(t *testing.T) {
	m, x := 10000, 100
	p := []int{1, 0, 1}
	runSample(t, m, x, p)
}

func TestSample3(t *testing.T) {
	m, x := 100000000, 1025
	p := []int{1, 0, 1, 0, 0, 0, 0}
	runSample(t, m, x, p)
}

func TestSample4(t *testing.T) {
	m, x := 30, 16
	p := []int{0, 1, 1, 1, 0}
	runSample(t, m, x, p)
}
