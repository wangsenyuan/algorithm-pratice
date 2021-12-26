package p2124

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, start []int, s string, expect []int) {
	res := executeInstructions(n, start, s)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	start := []int{0, 1}
	s := "RRDDLU"
	expect := []int{1, 5, 4, 3, 1, 0}
	runSample(t, n, start, s, expect)
}

func TestSample2(t *testing.T) {
	n := 1
	start := []int{0, 0}
	s := "LRUD"
	expect := []int{0, 0, 0, 0}
	runSample(t, n, start, s, expect)
}
