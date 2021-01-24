package p1734

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, encoded []int, expect []int) {
	res := decode(encoded)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	encoded := []int{3, 1}
	expect := []int{1, 2, 3}
	runSample(t, encoded, expect)
}

func TestSample2(t *testing.T) {
	encoded := []int{6, 5, 4, 6}
	expect := []int{2, 4, 1, 5, 3}
	runSample(t, encoded, expect)
}
