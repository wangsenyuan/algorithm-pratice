package main

import "testing"

func runSample(t *testing.T, a []int, expect []int) {
	res := solve(a)

	if res[0] != expect[0] {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	k := res[1]

	get := func(i int) int {
		if a[i] < 0 {
			return k
		}
		return a[i]
	}

	var m int
	for i := 1; i < len(a); i++ {
		v := get(i)
		w := get(i - 1)
		m = max(m, abs(v-w))
	}

	if m != res[0] {
		t.Fatalf("Sample result %v, not getting the correct answer", res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{-1, 10, -1, 12, -1}
	expect := []int{1, 11}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{-1, 40, 35, -1, 35}
	expect := []int{5, 35}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{-1, -1, 9, -1, 3, -1}
	expect := []int{3, 6}
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{-1, -1}
	expect := []int{0, 0}
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{0, -1}
	expect := []int{0, 0}
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{1, -1, 3, -1}
	expect := []int{1, 2}
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{1, -1, 7, 5, 2, -1, 5}
	expect := []int{3, 4}
	runSample(t, a, expect)
}
