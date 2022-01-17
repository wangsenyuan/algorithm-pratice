package p2141

import "testing"

func runSample(t *testing.T, n int, batteries []int, expect int) {
	res := int(maxRunTime(n, batteries))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	batteries := []int{3, 3, 3}
	expect := 4
	runSample(t, n, batteries, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	batteries := []int{1, 1, 1, 1}
	expect := 2
	runSample(t, n, batteries, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	batteries := []int{10, 10, 3, 5}
	expect := 8
	runSample(t, n, batteries, expect)
}
