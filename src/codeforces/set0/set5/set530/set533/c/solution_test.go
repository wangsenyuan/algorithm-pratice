package main

import "testing"

func runSample(t *testing.T, pos []int, expect string) {
	res := solve(pos)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	pos := []int{4, 7, 7, 4}
	expect := "Vasiliy"
	runSample(t, pos, expect)
}

func TestSample2(t *testing.T) {
	pos := []int{2, 1, 2, 2}
	expect := "Polycarp"
	runSample(t, pos, expect)
}

func TestSample3(t *testing.T) {
	pos := []int{0, 90000, 89999, 89999}
	expect := "Vasiliy"
	runSample(t, pos, expect)
}
