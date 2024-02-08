package main

import "testing"

func runSample(t *testing.T, a []int, expect []int) {
	res := solve(a)

	x := check(a, expect)
	y := check(a, res)

	if x != y {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func check(a []int, ans []int) int {
	var sum int
	prod := 1
	for i := 0; i < len(a); i++ {
		if i < ans[0] || i >= ans[1] {
			sum += a[i]
		} else {
			if prod > oo/a[i] {
				return oo
			}
			prod *= a[i]
		}
	}
	return sum + prod
}

func TestSample1(t *testing.T) {
	a := []int{1, 3, 1, 3}
	expect := []int{2, 4}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 2, 3}
	expect := []int{3, 4}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 1, 1, 1, 1}
	expect := []int{1, 1}
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{10, 1, 10, 1, 10}
	expect := []int{1, 5}
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{2, 1, 1, 3}
	expect := []int{4, 4}
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{2, 1, 2}
	expect := []int{1, 1}
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{1, 7}
	expect := []int{1, 1}
	runSample(t, a, expect)
}
