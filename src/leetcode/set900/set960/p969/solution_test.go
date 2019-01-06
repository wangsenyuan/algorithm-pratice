package p969

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, B []int) {
	res := pancakeSort(A)
	if !reflect.DeepEqual(A, B) {
		t.Errorf("Sample result %v, doesn't give correct answer, witch is %v", res, A)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{3, 2, 4, 1}, []int{1, 2, 3, 4})
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 2, 3}, []int{1, 2, 3})
}
