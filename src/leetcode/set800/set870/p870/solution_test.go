package p870

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, B []int, expect []int) {
	res := advantageCount(A, B)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 7, 11, 15}
	B := []int{1, 10, 4, 11}
	expect := []int{2, 11, 7, 15}
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{12, 24, 8, 32}
	B := []int{13, 25, 32, 11}
	expect := []int{24, 32, 8, 12}
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 0, 4, 1, 2}
	B := []int{1, 3, 0, 0, 2}
	expect := []int{2, 0, 2, 1, 4}
	runSample(t, A, B, expect)
}
