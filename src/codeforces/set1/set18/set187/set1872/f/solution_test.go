package main

import "testing"

func runSample(t *testing.T, c []int, a []int, expect []int) {
	res := solve(c, a)

	x := getMoney(c, a, expect)
	y := getMoney(c, a, res)

	if x != y {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func getMoney(c []int, a []int, ord []int) int {
	n := len(c)
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[ord[i]-1] = i
	}
	var res int
	for i := 0; i < n; i++ {
		j := a[i] - 1
		if pos[i] < pos[j] {
			res += 2 * c[i]
		} else {
			res += c[i]
		}
	}
	return res
}

func TestSample1(t *testing.T) {
	c := []int{6, 6, 1}
	a := []int{2, 3, 2}
	expect := []int{1, 2, 3}
	runSample(t, c, a, expect)
}

func TestSample2(t *testing.T) {
	c := []int{1, 2, 1, 2, 2, 1, 2, 1}
	a := []int{2, 1, 4, 3, 6, 5, 8, 7}
	expect := []int{2, 4, 5, 1, 6, 3, 7, 8}
	runSample(t, c, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 1, 1, 1, 1}
	c := []int{9, 8, 1, 1, 1}
	expect := []int{3, 4, 5, 1, 2}
	runSample(t, c, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{2, 1}
	c := []int{1000000000, 999999999}
	expect := []int{1, 2}
	runSample(t, c, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{2, 3, 2, 6, 4, 4, 3}
	c := []int{1, 2, 3, 4, 5, 6, 7}
	expect := []int{7, 5, 1, 3, 2, 6, 4}
	runSample(t, c, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{3, 4, 4, 1, 3}
	c := []int{3, 4, 5, 6, 7}
	expect := []int{5, 2, 3, 4, 1}
	runSample(t, c, a, expect)
}
