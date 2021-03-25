package p082

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, arr []int, expect []int) {
	head := NewList(arr)
	res := deleteDuplicates(head)
	ans := res.ToArray()

	if !reflect.DeepEqual(expect, ans) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 2, 3, 3, 4, 4, 5}
	expect := []int{1, 2, 5}
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{1, 1, 1, 2, 3}
	expect := []int{2, 3}
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1, 1, 2, 2, 3}
	expect := []int{3}
	runSample(t, arr, expect)
}
