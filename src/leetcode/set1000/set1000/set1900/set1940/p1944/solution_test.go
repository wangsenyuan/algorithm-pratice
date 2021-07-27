package p1944

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, heights []int, expect []int) {
	res := canSeePersonsCount(heights)

	if !reflect.DeepEqual(expect, res) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	heights := []int{10, 6, 8, 5, 11, 9}
	expect := []int{3, 1, 2, 1, 1, 0}
	runSample(t, heights, expect)
}

func TestSample2(t *testing.T) {
	heights := []int{5, 1, 2, 3, 10}
	expect := []int{4, 1, 1, 1, 0}
	runSample(t, heights, expect)
}
