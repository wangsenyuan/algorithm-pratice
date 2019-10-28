package p1238

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, start int, expect []int) {
	res := circularPermutation(n, start)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d, expect %v, but got %v", n, start, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, []int{3, 2, 0, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 2, []int{2, 6, 7, 5, 4, 0, 1, 3})
}
