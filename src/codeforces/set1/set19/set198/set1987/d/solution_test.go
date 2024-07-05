package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 4, 2, 3}
	expect := 2
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 1}
	expect := 1
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 4, 2, 3, 4}
	expect := 3
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{3, 4, 1, 4}
	expect := 2
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{1}
	expect := 1
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{4, 3, 2, 5, 6, 8, 3, 4}
	// bob处聊掉5,6,8后，alice只能得到3个
	expect := 3
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{6, 1, 1, 3, 5, 3, 1}
	expect := 2
	runSample(t, a, expect)
}

func TestSample8(t *testing.T) {
	a := []int{6, 11, 6, 8, 7, 5, 3, 11, 2, 3, 5}
	expect := 4
	runSample(t, a, expect)
}

func TestSample9(t *testing.T) {
	a := []int{2, 6, 5, 3, 9, 1, 6, 2, 5, 6, 3, 2, 3, 9, 6, 1, 6}
	expect := 4
	runSample(t, a, expect)
}
