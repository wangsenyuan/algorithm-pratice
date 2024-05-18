package main

import "testing"

func runSample(t *testing.T, c int, s []int, expect int) {
	res := solve(c, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	c := 3
	s := []int{1, 2, 3}
	expect := 3
	runSample(t, c, s, expect)
}

func TestSample2(t *testing.T) {
	c := 179
	s := []int{57}
	expect := 16139
	runSample(t, c, s, expect)
}

func TestSample3(t *testing.T) {
	c := 6
	s := []int{0, 3, 5, 6}
	expect := 10
	runSample(t, c, s, expect)
}

func TestSample4(t *testing.T) {
	c := 1
	s := []int{1}
	expect := 2
	runSample(t, c, s, expect)
}

func TestSample5(t *testing.T) {
	c := 10
	s := []int{0, 2, 4, 8, 10}
	expect := 33
	runSample(t, c, s, expect)
}

func TestSample6(t *testing.T) {
	c := 10
	s := []int{1, 3, 5, 7, 9}
	expect := 36
	runSample(t, c, s, expect)
}

func TestSample7(t *testing.T) {
	c := 10
	s := []int{2, 4, 6, 7}
	expect := 35
	runSample(t, c, s, expect)
}

func TestSample8(t *testing.T) {
	c := 1000000000
	s := []int{228, 1337, 998244353}
	expect := 499999998999122959
	runSample(t, c, s, expect)
}
