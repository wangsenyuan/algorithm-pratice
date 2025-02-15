package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 4
4 3 7 5
2 5 4 6`
	expect := []int{7, 4, 2, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 5
5 0 4 8 6
1 3 9 2 7`
	expect := []int{6, 2, 1, 1, 2}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 7
1 1 4 5 1 4
1 9 1 9 8 1 0`
	expect := []int{7, 3, 2, 1, 1, 1, 1}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `7 6
1 9 1 9 8 1 0
1 1 4 5 1 4`
	expect := []int{15, 9, 5, 4, 4, 4}
	runSample(t, s, expect)
}
