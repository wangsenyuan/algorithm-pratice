package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
BGBG
4 2 4 3
`
	expect := [][]int{
		{3, 4},
		{1, 2},
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
BBGG
4 6 1 5
`
	expect := [][]int{
		{2, 3},
		{1, 4},
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
BGBB
1 1 2 3
`
	expect := [][]int{
		{1, 2},
	}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10
BGGBBGBGBG
9892816 3514007 5425956 5241945 9171176 3351177 2772494 2891569 1510552 8471969
`
	expect := [][]int{
		{7, 8},
		{3, 4},
		{6, 9},
		{5, 10},
		{1, 2},
	}
	runSample(t, s, expect)
}
