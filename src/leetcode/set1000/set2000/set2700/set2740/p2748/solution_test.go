package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, pos []int, hp []int, dir string, expect []int) {
	res := survivedRobotsHealths(pos, hp, dir)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	positions := []int{5, 4, 3, 2, 1}
	healths := []int{2, 17, 9, 15, 10}
	directions := "RRRRR"
	expect := []int{2, 17, 9, 15, 10}
	runSample(t, positions, healths, directions, expect)
}

func TestSample2(t *testing.T) {
	positions := []int{3, 5, 2, 6}
	healths := []int{10, 10, 15, 12}
	directions := "RLRL"
	expect := []int{14}
	runSample(t, positions, healths, directions, expect)
}

func TestSample3(t *testing.T) {
	positions := []int{1, 2, 5, 6}
	healths := []int{10, 10, 11, 11}
	directions := "RLRL"
	runSample(t, positions, healths, directions, nil)
}
