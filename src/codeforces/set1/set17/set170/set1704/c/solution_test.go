package main

import "testing"

func runSample(t *testing.T, n int, m int, infected []int, expect int) {
	res := solve(n, m, infected)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 10, 3
	infected := []int{3, 6, 8}
	expect := 7
	runSample(t, n, m, infected, expect)
}

func TestSample2(t *testing.T) {
	n, m := 6, 2
	infected := []int{2, 5}
	expect := 5
	runSample(t, n, m, infected, expect)
}

func TestSample3(t *testing.T) {
	n, m := 20, 3
	infected := []int{3, 7, 12}
	expect := 11
	runSample(t, n, m, infected, expect)
}

func TestSample4(t *testing.T) {
	n, m := 41, 5
	infected := []int{1, 11, 21, 31, 41}
	expect := 28
	runSample(t, n, m, infected, expect)
}

func TestSample5(t *testing.T) {
	n, m := 10, 5
	infected := []int{2, 4, 6, 8, 10}
	expect := 9
	runSample(t, n, m, infected, expect)
}

func TestSample6(t *testing.T) {
	n, m := 5, 5
	infected := []int{3, 2, 5, 4, 1}
	expect := 5
	runSample(t, n, m, infected, expect)
}

func TestSample7(t *testing.T) {
	n, m := 1000000000, 1
	infected := []int{1}
	expect := 2
	runSample(t, n, m, infected, expect)
}

func TestSample8(t *testing.T) {
	n, m := 1000000000, 4
	infected := []int{1, 1000000000, 10, 16}
	expect := 15
	runSample(t, n, m, infected, expect)
}
