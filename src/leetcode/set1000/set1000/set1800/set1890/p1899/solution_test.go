package p1899

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, first, second int, expect []int) {
	res := earliestAndLatest(n, first, second)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 11
	first := 2
	second := 4
	expect := []int{3, 4}
	runSample(t, n, first, second, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	first := 1
	second := 5
	expect := []int{1, 1}
	runSample(t, n, first, second, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	first := 1
	second := 4
	expect := []int{2, 2}
	runSample(t, n, first, second, expect)
}
