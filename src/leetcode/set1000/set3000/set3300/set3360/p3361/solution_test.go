package p3361

import "testing"

func runSample(t *testing.T, s string, x string, next []int, prev []int, expect int) {
	res := shiftDistance(s, x, next, prev)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abab"
	x := "baba"
	next := []int{100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	prev := []int{1, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	expect := 2
	runSample(t, s, x, next, prev, expect)
}

func TestSample2(t *testing.T) {
	s := "leet"
	x := "code"
	next := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	prev := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	expect := 31
	runSample(t, s, x, next, prev, expect)
}
