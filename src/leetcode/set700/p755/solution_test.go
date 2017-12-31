package p755

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, heights []int, V int, K int, expect []int) {
	res := pourWater(heights, V, K)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample: %v %d %d, should give %v, but got %v", heights, V, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	heights := []int{2, 1, 1, 2, 1, 2, 2}
	V, K := 4, 3
	expect := []int{2, 2, 2, 3, 2, 2, 2}
	runSample(t, heights, V, K, expect)
}

func TestSample2(t *testing.T) {
	heights := []int{1, 2, 3, 4}
	V, K := 2, 2
	expect := []int{2, 3, 3, 4}
	runSample(t, heights, V, K, expect)
}

func TestSample3(t *testing.T) {
	heights := []int{3, 1, 3}
	V, K := 5, 1
	expect := []int{4, 4, 4}
	runSample(t, heights, V, K, expect)
}
