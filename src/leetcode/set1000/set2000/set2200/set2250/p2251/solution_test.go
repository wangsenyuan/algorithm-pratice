package p2251

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, flowers [][]int, persons []int, expect []int) {
	res := fullBloomFlowers(flowers, persons)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	flowers := [][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}}
	persons := []int{2, 3, 7, 11}
	expect := []int{1, 2, 2, 2}
	runSample(t, flowers, persons, expect)
}

func TestSample2(t *testing.T) {
	flowers := [][]int{{1, 10}, {3, 3}}
	persons := []int{3, 3, 2}
	expect := []int{2, 2, 1}
	runSample(t, flowers, persons, expect)
}
