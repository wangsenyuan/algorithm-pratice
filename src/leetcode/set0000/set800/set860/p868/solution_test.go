package p868

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A [][]int, expect [][]int) {
	res := transpose(A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}
	expect := [][]int{
		{1, 4, 7}, {2, 5, 8}, {3, 6, 9},
	}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := [][]int{
		{1, 2, 3}, {4, 5, 6},
	}
	expect := [][]int{
		{1, 4}, {2, 5}, {3, 6},
	}
	runSample(t, A, expect)
}
