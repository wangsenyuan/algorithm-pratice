package p1073

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, arr1 []int, arr2 []int, expect []int) {
	res := addNegabinary(arr1, arr2)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", arr1, arr2, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr1 := []int{1, 1, 1, 1, 1}
	arr2 := []int{1, 0, 1}
	expect := []int{1, 0, 0, 0, 0}
	runSample(t, arr1, arr2, expect)
}
