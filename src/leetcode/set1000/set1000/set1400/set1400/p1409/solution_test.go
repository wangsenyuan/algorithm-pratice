package p1409

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, queries []int, m int, expect []int) {
	res := processQueries(queries, m)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample fail, expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	queries := []int{3, 1, 2, 1}
	m := 5
	expect := []int{2, 1, 2, 1}
	runSample(t, queries, m, expect)
}

func TestSample2(t *testing.T) {
	queries := []int{4, 1, 2, 2}
	m := 4
	expect := []int{3, 1, 2, 0}
	runSample(t, queries, m, expect)
}

func TestSample3(t *testing.T) {
	queries := []int{7, 5, 5, 8, 3}
	m := 8
	expect := []int{6, 5, 0, 7, 5}
	runSample(t, queries, m, expect)
}
