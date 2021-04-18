package p1828

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, points [][]int, queries [][]int, expect []int) {
	res := countPoints(points, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ps := [][]int{
		{1, 3}, {3, 3}, {5, 3}, {2, 2},
	}
	qs := [][]int{
		{2, 3, 1}, {4, 3, 1}, {1, 1, 2},
	}
	expect := []int{3, 2, 2}
	runSample(t, ps, qs, expect)
}

func TestSample2(t *testing.T) {
	ps := [][]int{
		{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5},
	}
	qs := [][]int{
		{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3},
	}
	expect := []int{2, 3, 2, 4}
	runSample(t, ps, qs, expect)
}

func TestSample3(t *testing.T) {
	ps := [][]int{
		{99, 113}, {150, 165}, {23, 65}, {175, 154}, {84, 83}, {24, 59}, {124, 29}, {19, 97}, {117, 182}, {105, 191}, {83, 117}, {114, 35}, {0, 111}, {22, 53},
	}
	qs := [][]int{
		{105, 191, 155}, {114, 35, 94}, {84, 83, 68}, {175, 154, 28}, {99, 113, 80}, {175, 154, 177}, {175, 154, 181}, {114, 35, 134}, {22, 53, 105}, {124, 29, 164}, {6, 99, 39}, {84, 83, 35},
	}
	expect := []int{11, 7, 8, 2, 7, 11, 13, 10, 10, 14, 3, 3}
	runSample(t, ps, qs, expect)
}
