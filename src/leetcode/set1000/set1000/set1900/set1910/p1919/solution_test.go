package p1919

import "testing"

func runSample(t *testing.T, m, n int, expect int) {
	res := colorTheGrid(m, n)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, n := 1, 1
	expect := 3
	runSample(t, m, n, expect)
}

func TestSample2(t *testing.T) {
	m, n := 1, 2
	expect := 6
	runSample(t, m, n, expect)
}

func TestSample3(t *testing.T) {
	m, n := 5, 5
	expect := 580986
	runSample(t, m, n, expect)
}
