package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 3, 6, 12, 17}
	expect := 24
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{6, 12, 8, 10, 15, 12, 18, 16}
	expect := 203
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{6, 12, 8}
	expect := bruteForce(a)
	runSample(t, a, expect)
}

func bruteForce(a []int) int {
	f := func(i, j, k int) int {
		x := min(a[i], a[j], a[k])
		y := a[i] + a[j] + a[k] - x - max(a[i], a[j], a[k])
		return gcd(x, y)
	}

	var res int
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			for k := j + 1; k < len(a); k++ {
				res += f(i, j, k)
			}
		}
	}

	return res
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
