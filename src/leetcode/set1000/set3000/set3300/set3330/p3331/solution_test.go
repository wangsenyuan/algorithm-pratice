package p3331

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, parent []int, s string, expect []int) {
	res := findSubtreeSizes(parent, s)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

}

func TestSample1(t *testing.T) {
	parent := []int{-1, 0, 0, 1, 1, 1}
	s := "abaabc"
	expect := []int{6, 3, 1, 1, 1, 1}
	runSample(t, parent, s, expect)
}
