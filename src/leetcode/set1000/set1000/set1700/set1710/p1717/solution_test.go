package p1717

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, expect []int) {
	res := constructDistancedSequence(n)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	expect := []int{5, 3, 1, 4, 3, 5, 2, 4, 2}
	runSample(t, n, expect)
}
