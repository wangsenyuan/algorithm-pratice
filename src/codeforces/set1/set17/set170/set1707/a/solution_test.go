package main

import "testing"

func runSample(t *testing.T, A []int, iq int, expect int) {
	res := solve(A, iq)

	var cnt int
	for i := 0; i < len(res); i++ {
		cnt += int(res[i] - '0')
	}

	if cnt != expect {
		t.Errorf("Sample expect %d, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1}
	iq := 1
	expect := 1
	runSample(t, A, iq, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2}
	iq := 1
	expect := 2
	runSample(t, A, iq, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 1}
	iq := 1
	expect := 2
	runSample(t, A, iq, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 4, 3, 1}
	iq := 2
	expect := 3
	runSample(t, A, iq, expect)
}

func TestSample5(t *testing.T) {
	A := []int{5, 1, 2, 4, 3}
	iq := 2
	expect := 4
	runSample(t, A, iq, expect)
}
