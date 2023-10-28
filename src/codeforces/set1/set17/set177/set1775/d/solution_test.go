package main

import "testing"

func runSample(t *testing.T, a []int, s int, x int, expect int) {
	res := solve(a, s, x)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}

	if expect == 0 {
		return
	}

	m := len(res)
	if res[0] != s || res[m-1] != x {
		t.Fatalf("Sample result %v, not correct", res)
	}

	for i := 0; i+1 < m; i++ {
		tmp := gcd(a[res[i]-1], a[res[i+1]-1])
		if tmp == 1 {
			t.Fatalf("Sample result %v not correct", res)
		}
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func TestSample1(t *testing.T) {
	a := []int{2, 14, 9, 6, 8, 15, 11}
	s, x := 5, 6
	expect := 3
	runSample(t, a, s, x, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 14, 9, 6, 8, 15, 11}
	s, x := 5, 7
	expect := 0
	runSample(t, a, s, x, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 14, 9, 6, 8, 15, 11}
	s, x := 5, 5
	expect := 1
	runSample(t, a, s, x, expect)
}
