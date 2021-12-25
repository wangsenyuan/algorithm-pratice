package main

import (
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)
	if res != expect {
		t.Errorf("sample %d %v, expect %d, but got %d", n, A, expect,res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 3, 4, 5 ,6}
	expect := 14
	runSample(t, 5, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{7, 7, 7, 7}
	expect := 0
	runSample(t, 4, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 1, 3}
	expect := 2
	runSample(t, 3, A, expect)
}
