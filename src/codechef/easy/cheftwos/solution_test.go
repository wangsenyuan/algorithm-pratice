package main

import "testing"

func runSample(t *testing.T, n int, a, b string, expect int) {
	res := solve(n, []byte(a), []byte(b))

	if res != expect {
		t.Errorf("sample %d %s %s, expect %d, but got %d", n, a, b, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	a, b := "1111", "2211"
	expect := 8
	runSample(t, n, a, b, expect)
}
func TestSample2(t *testing.T) {
	n := 3
	a, b := "222", "111"
	expect := 0
	runSample(t, n, a, b, expect)
}
