package p054

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, matrix [][]int, expect []int) {
	res := spiralOrder(matrix)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %v, should give %v, but got %v", matrix, expect, res)
	}
}

func TestSample1(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expect := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	runSample(t, matrix, expect)
}
