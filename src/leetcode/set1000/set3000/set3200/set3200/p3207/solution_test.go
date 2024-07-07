package p3207

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := numberOfAlternatingGroups(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{0, 1, 0, 1, 0}
	k := 3
	expect := 3
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{0, 1, 0, 0, 1, 0, 1}
	k := 6
	expect := 2
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 1, 0, 1}
	k := 4
	expect := 0
	runSample(t, a, k, expect)
}

func TestSample4(t *testing.T) {
	a := []int{0, 0, 0}
	k := 3
	expect := 0
	runSample(t, a, k, expect)
}