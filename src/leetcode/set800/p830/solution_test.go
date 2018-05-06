package p830

import (
	"testing"
	"reflect"
)

func TestSample(t *testing.T) {
	S := "abbxxxxzzyyy"
	res := largeGroupPositions(S)
	expect := [][]int{{3, 6}, {9, 11}}
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("error, got %v", res)
	}
}
