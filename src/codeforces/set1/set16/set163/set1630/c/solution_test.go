package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 1, 2, 7, 4, 7}
	expect := 2
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 2, 1, 3, 3, 4, 5, 5, 5, 4, 7}
	expect := 7
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3, 4, 1, 3, 4}
	expect := 4
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 2, 3, 4, 1, 5, 7, 5, 3}
	expect := 6
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{1, 2, 3, 4, 1, 5, 7, 3, 5}
	// 为啥前面那个5不能被处理呢？
	expect := 5
	runSample(t, A, expect)
}

func TestSample6(t *testing.T) {
	A := []int{1, 2, 3, 4, 1, 6, 5, 6, 7, 3, 5}

	expect := 7
	runSample(t, A, expect)
}
