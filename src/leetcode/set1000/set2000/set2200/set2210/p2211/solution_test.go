package p2211

import "testing"

func runSample(t *testing.T, dir string, expect int) {
	res := countCollisions(dir)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	dir := "RLRSLL"
	expect := 5
	runSample(t, dir, expect)
}

func TestSample2(t *testing.T) {
	dir := "LLRR"
	expect := 0
	runSample(t, dir, expect)
}
