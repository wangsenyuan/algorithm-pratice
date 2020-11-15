package p5604

import "testing"

func runSample(t *testing.T, m int, n int, in int, out int, expect int) {
	res := getMaxGridHappiness(m, n, in, out)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 2
	n := 3
	introvertsCount := 1
	extrovertsCount := 2
	expect := 240
	runSample(t, m, n, introvertsCount, extrovertsCount, expect)
}

func TestSample2(t *testing.T) {
	m := 3
	n := 1
	introvertsCount := 2
	extrovertsCount := 1
	expect := 260
	runSample(t, m, n, introvertsCount, extrovertsCount, expect)
}
