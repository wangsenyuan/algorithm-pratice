package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, queries [][]int, expect []int) {
	res := solve(a, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{15, 14, 17, 42, 34}
	queries := [][]int{
		{1, 7},
		{2, 15},
		{4, 5},
	}
	expect := []int{2, -1, 5}

	runSample(t, a, queries, expect)
}

func TestSample2(t *testing.T) {
	a := []int{7, 5, 3, 1, 7}
	queries := [][]int{
		{1, 7},
		{5, 7},
		{2, 3},
		{2, 2},
	}
	expect := []int{1, 5, 2, 2}

	runSample(t, a, queries, expect)
}

func TestSample3(t *testing.T) {
	a := []int{19, 20, 15, 12, 21, 7, 11}
	queries := [][]int{
		{1, 15},
		{4, 4},
		{7, 12},
		{5, 7},
	}
	expect := []int{2, 6, -1, 5}

	runSample(t, a, queries, expect)
}

func TestSample4(t *testing.T) {
	a := []int{321743262, 571897166, 908254205, 974719387, 163128565, 87201363, 825121690, 373710421, 201168977, 260453815, 755664311}
	queries := [][]int{
		{11, 601796739},
		{4, 953622832},
		{5, 515176969},
		{9, 941890832},
		{6, 219905114},
		{6, 427059990},
	}
	expect := []int{11, 4, -1, -1, -1, -1}

	runSample(t, a, queries, expect)
}
