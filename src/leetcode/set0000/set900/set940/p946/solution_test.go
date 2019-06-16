package p946

import "testing"

func runSample(t *testing.T, pushed, popped []int, expect bool) {
	res := validateStackSequences(pushed, popped)
	if res != expect {
		t.Errorf("Sample %v, %v, expect %t, but got %t", pushed, popped, expect, res)
	}
}

func TestSample1(t *testing.T) {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	expect := true
	runSample(t, pushed, popped, expect)
}

func TestSample2(t *testing.T) {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 3, 5, 1, 2}
	expect := false
	runSample(t, pushed, popped, expect)
}

func TestSample3(t *testing.T) {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{5, 4, 3, 2, 1}
	expect := true
	runSample(t, pushed, popped, expect)
}
