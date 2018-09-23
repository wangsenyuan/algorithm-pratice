package p911

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, persons []int, times []int, queries []int, expect []int) {
	sol := Constructor(persons, times)

	res := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		res[i] = sol.Q(queries[i])

	}
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v %v, expect %v, but got %v", persons, times, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	persons := []int{0, 1, 1, 0, 0, 1, 0}
	times := []int{0, 5, 10, 15, 20, 25, 30}
	queries := []int{3, 12, 25, 15, 4, 8}
	expect := []int{0, 1, 1, 0, 0, 1}
	runSample(t, persons, times, queries, expect)
}
