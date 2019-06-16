package main

import "testing"

func TestFindMaxForm1(t *testing.T) {
	strs := []string{"111", "1000", "1000", "1000"}
	m := 9
	n := 3
	expect := 3
	got := findMaxForm(strs, m, n)
	if got != expect {
		t.Errorf("%v, %d, %d should get %d, but get %d\n", strs, m, n, expect, got)
	}
}
