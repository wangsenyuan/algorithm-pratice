package main

import "testing"

func runSample(t *testing.T, teachers []int, groups [][]int, expect string) {
	res := solve(teachers, groups)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	teachers := []int{30}
	groups := [][]int{{25, 16, 37}}
	expect := "101"
	runSample(t, teachers, groups, expect)
}

func TestSample2(t *testing.T) {
	teachers := []int{9, 12, 12, 6}
	groups := [][]int{{4, 5}, {111, 11, 11}}
	expect := "00100"
	runSample(t, teachers, groups, expect)
}
