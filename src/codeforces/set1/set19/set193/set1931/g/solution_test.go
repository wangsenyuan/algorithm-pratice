package main

import "testing"

func runSample(t *testing.T, cnt []int, expect int) {
	res := solve(cnt)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cnt := []int{1, 1, 1, 1}
	expect := 4
	runSample(t, cnt, expect)
}

func TestSample2(t *testing.T) {
	cnt := []int{1, 2, 5, 10}
	expect := 66
	runSample(t, cnt, expect)
}

func TestSample3(t *testing.T) {
	cnt := []int{900000, 900000, 900000, 900000}
	expect := 794100779
	runSample(t, cnt, expect)
}

func TestSample4(t *testing.T) {
	cnt := []int{0, 0, 0, 4}
	expect := 1
	runSample(t, cnt, expect)
}
