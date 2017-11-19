package p728

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}
	left, right := 1, 22
	res := selfDividingNumbers(left, right)
	if !reflect.DeepEqual(expect, res) {
		t.Errorf("self dividing numbers between %d and %d should be %v, but got %v", left, right, expect, res)
	}
}
