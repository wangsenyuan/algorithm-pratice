package p1613

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, k int, arrival []int, load []int, expect []int) {
	res := busiestServers(k, arrival, load)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 3
	arrival := []int{1, 2, 3, 4, 5}
	load := []int{5, 2, 3, 3, 3}
	expect := []int{1}
	runSample(t, k, arrival, load, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	arrival := []int{1, 2, 3, 4}
	load := []int{1, 2, 1, 2}
	expect := []int{0}
	runSample(t, k, arrival, load, expect)
}

func TestSample3(t *testing.T) {
	k := 3
	arrival := []int{1, 2, 3}
	load := []int{10, 12, 11}
	expect := []int{0, 1, 2}
	runSample(t, k, arrival, load, expect)
}

func TestSample4(t *testing.T) {
	k := 3
	arrival := []int{1, 2, 3, 4, 8, 9, 10}
	load := []int{5, 2, 10, 3, 1, 2, 2}
	expect := []int{1}
	runSample(t, k, arrival, load, expect)
}

func TestSample5(t *testing.T) {
	k := 1
	arrival := []int{1}
	load := []int{1}
	expect := []int{0}
	runSample(t, k, arrival, load, expect)
}
