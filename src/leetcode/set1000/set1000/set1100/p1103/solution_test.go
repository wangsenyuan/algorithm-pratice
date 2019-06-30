package p1103

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, candies int, n int, expect []int) {
	res := distributeCandies(candies, n)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %d, expect %v, but got %v", candies, n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	expect := []int{1, 2, 3, 1}
	runSample(t, 7, 4, expect)
}

func TestSample2(t *testing.T) {
	expect := []int{5, 2, 3}
	runSample(t, 10, 3, expect)
}

func TestSample3(t *testing.T) {
	expect := []int{4, 6}
	runSample(t, 10, 2, expect)
}

func TestSample4(t *testing.T) {
	expect := []int{6, 6}
	runSample(t, 12, 2, expect)
}
