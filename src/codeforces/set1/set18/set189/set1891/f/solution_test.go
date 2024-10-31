package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, ops [][]int, expect []int) {
	res := solve(ops)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ops := [][]int{
		{2, 1, 3},
		{1, 1},
		{2, 2, 1},
		{1, 1},
		{2, 3, 2},
		{1, 3},
		{2, 1, 4},
		{1, 3},
		{2, 3, 2},
	}
	expect := []int{7, 5, 8, 6, 2}

	runSample(t, ops, expect)
}

func TestSample2(t *testing.T) {
	ops := [][]int{
		{2, 1, 1},
		{1, 1},
		{2, 1, -1},
		{1, 1},
		{2, 1, 1},
	}
	expect := []int{1, 0, 1}

	runSample(t, ops, expect)
}

func TestSample3(t *testing.T) {
	ops := [][]int{
		{1, 1},
		{1, 1},
		{2, 1, 1},
		{2, 1, 3},
		{2, 2, 10},
	}
	expect := []int{4, 14, 4}

	runSample(t, ops, expect)
}

func TestSample4(t *testing.T) {
	ops := [][]int{
		{2, 1, -237652274},
		{2, 1, 355837696},
		{2, 1, 40127609},
		{2, 1, -951430355},
		{2, 1, -952998006},
		{2, 1, 197332879},
		{2, 1, -705358756},
		{2, 1, 494390475},
		{2, 1, 174099103},
		{2, 1, 39291469},
		{2, 1, -704876538},
		{2, 1, -970819209},
		{2, 1, -564313440},
		{2, 1, -108380022},
		{2, 1, 287260804},
		{2, 1, -748427729},
		{2, 1, 689176267},
		{2, 1, -508648938},
		{2, 1, 535513443},
		{2, 1, 995953274},
		{2, 1, -956770927},
		{2, 1, 692400840},
		{2, 1, 251138894},
		{2, 1, 425166489},
		{2, 1, 206100593},
		{2, 1, -802403316},
		{2, 1, -4762106},
		{2, 1, -168989533},
		{2, 1, -774000240},
		{2, 1, 108313813},
		{2, 1, 839921068},
		{2, 1, -58159178},
	}
	expect := []int{4, 14, 4}

	runSample(t, ops, expect)
}
