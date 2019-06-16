package p735

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	samples := []struct {
		asteroids []int
		result    []int
	}{
		{[]int{5, 10, -5}, []int{5, 10}},
		{[]int{8, -8}, []int{}},
		{[]int{10, 2, -5}, []int{10}},
		{[]int{-2, -1, 1, 2}, []int{-2, -1, 1, 2}},
	}

	for _, sample := range samples {
		res := asteroidCollision(sample.asteroids)
		if !reflect.DeepEqual(res, sample.result) {
			t.Errorf("sample %v, after collision, should leave %v, but got %v", sample.asteroids, sample.result, res)
		}
	}
}
