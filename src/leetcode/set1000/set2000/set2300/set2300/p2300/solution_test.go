package p2300

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, spells []int, potions []int, success int64, expect []int) {
	res := successfulPairs(spells, potions, success)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	spells := []int{5, 1, 3}
	potions := []int{1, 2, 3, 4, 5}
	var success int64 = 7
	expect := []int{4, 0, 3}
	runSample(t, spells, potions, success, expect)
}

func TestSample2(t *testing.T) {
	spells := []int{3, 1, 2}
	potions := []int{8, 5, 8}
	var success int64 = 16
	expect := []int{2, 0, 2}
	runSample(t, spells, potions, success, expect)
}
