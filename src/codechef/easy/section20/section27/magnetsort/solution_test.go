package main

import "testing"

func runSample(t *testing.T, n int, A []int, polarity string, expect int) {
	res := solve(n, A, polarity)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 3, 2, 3, 7}
	polarity := "NNSNS"
	expect := 1
	runSample(t, n, A, polarity, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []int{2, 1}
	polarity := "SS"
	expect := -1
	runSample(t, n, A, polarity, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	A := []int{1, 8, 12}
	polarity := "SNS"
	expect := 0
	runSample(t, n, A, polarity, expect)
}
func TestSample4(t *testing.T) {
	n := 3
	A := []int{3, 2, 1}
	polarity := "NSN"
	expect := 2
	runSample(t, n, A, polarity, expect)
}

func TestSample5(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 5, 4}
	polarity := "NSNSN"
	expect := 1
	runSample(t, n, A, polarity, expect)
}

func TestSample6(t *testing.T) {
	n := 5
	A := []int{1, 1, 2, 2, 1}
	polarity := "SNSNN"
	expect := 1
	runSample(t, n, A, polarity, expect)
}
