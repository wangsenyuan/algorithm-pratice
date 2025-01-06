package main

import "testing"

func runSample(t *testing.T, s string, expect []int) {
	res := solve(s)

	get := func(res []int) string {
		i, j, u, v := res[0], res[1], res[2], res[3]
		x := s[i-1 : j]
		y := s[u-1 : v]
		n := len(x)
		m := len(y)
		k := max(m, n)
		z := make([]byte, k)

		i = 0
		for i < min(m, n) {
			if x[n-1-i] == y[m-1-i] {
				z[k-1-i] = '0'
			} else {
				z[k-1-i] = '1'
			}
			i++
		}
		for i < m {
			z[k-1-i] = y[m-1-i]
			i++
		}
		for i < n {
			z[k-1-i] = x[n-1-i]
			i++
		}
		return string(z)
	}

	a := get(expect)
	b := get(res)

	if a != b {
		t.Fatalf("Sample result %v, not correct, it get %s, instead of %s", res, b, a)
	}
}

func TestSample1(t *testing.T) {
	s := "111"
	expect := []int{2, 2, 1, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "1000"
	expect := []int{1, 3, 1, 4}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "10111"
	expect := []int{1, 5, 1, 4}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "11101"
	expect := []int{3, 4, 1, 5}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "1100010001101"
	expect := []int{1, 13, 1, 11}
	runSample(t, s, expect)
}
