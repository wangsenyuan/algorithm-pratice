package p1177

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, quries [][]int, expect []bool) {
	res := canMakePaliQueries(s, quries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %s %v, expect %v, but got %v", s, quries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcda"
	queries := [][]int{
		{3, 3, 0}, {1, 2, 0}, {0, 3, 1}, {0, 3, 2}, {0, 4, 1},
	}
	expect := []bool{true, false, false, true, true}
	runSample(t, s, queries, expect)
}
