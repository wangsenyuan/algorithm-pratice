package main

import "testing"

func runSample(t *testing.T, a, n, m int64, expect int64) {
	res := solve(a, n, m)
	if res != expect {
		t.Errorf("sample %d %d %d, should give %d, but got %d", a, n, m, expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, n, m := int64(12), int64(2), int64(17)
	expect := int64(5)
	runSample(t, a, n, m, expect)
}

func TestSample2(t *testing.T) {
	a, n, m := int64(523), int64(3), int64(11)
	expect := int64(6)
	runSample(t, a, n, m, expect)
}
