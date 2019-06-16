package p733

import (
	"reflect"
	"testing"
)

func TestSample1(t *testing.T) {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	expect := [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}
	res := floodFill(image, 1, 1, 2)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("flood fill on the sample %v should get %v, but get %v", image, expect, res)
	}
}

func TestSample2(t *testing.T) {
	image := [][]int{{0, 0, 0}, {0, 1, 0}}
	expect := [][]int{{2, 2, 2}, {2, 1, 2}}
	res := floodFill(image, 1, 0, 2)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("flood fill on the sample %v should get %v, but get %v", image, expect, res)
	}
}

func TestSample3(t *testing.T) {
	image := [][]int{{0, 0, 0}, {0, 0, 0}}
	expect := [][]int{{2, 2, 2}, {2, 2, 2}}
	res := floodFill(image, 0, 0, 2)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("flood fill on the sample %v should get %v, but get %v", image, expect, res)
	}
}
