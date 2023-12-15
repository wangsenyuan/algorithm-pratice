package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d,but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 7, 7, 2, 7}
	expect := 6
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{6, 2, 18}
	expect := 1
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := 3
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1000, 993, 986, 179}
	expect := 0
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	expect := 9
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{1, 1, 2, 2, 4, 4, 8, 8}
	expect := 16
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{1, 1, 1, 2, 2, 2, 4, 4, 4}
	expect := 45
	runSample(t, a, expect)
}
