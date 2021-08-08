package p1964

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, obs []int, expect []int) {
	res := longestObstacleCourseAtEachPosition(obs)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	obs := []int{1, 2, 3, 2}
	expect := []int{1, 2, 3, 3}
	runSample(t, obs, expect)
}

func TestSample2(t *testing.T) {
	obs := []int{2, 2, 1}
	expect := []int{1, 2, 1}
	runSample(t, obs, expect)
}

func TestSample3(t *testing.T) {
	obs := []int{3, 1, 5, 6, 4, 2}
	expect := []int{1, 1, 2, 3, 2, 2}
	runSample(t, obs, expect)
}
