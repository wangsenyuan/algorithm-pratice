package p927

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, expect []int) {
	res := threeEqualParts(A)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 0, 1, 0, 1}, []int{0, 3})
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 1, 0, 1, 1}, []int{-1, -1})
}

func TestSample3(t *testing.T) {
	runSample(t, []int{0, 0, 0, 0, 0}, []int{0, 2})
}

func TestSample4(t *testing.T) {
	runSample(t, []int{1, 0, 1, 1, 0}, []int{-1, -1})
}
