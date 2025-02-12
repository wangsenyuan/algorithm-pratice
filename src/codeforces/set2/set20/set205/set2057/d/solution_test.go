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
	s := `2 2
1 10
1 10
2 2`
	expect := []int{8, 0, 7}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 3
1 2 3 4 5
3 7
1 4
5 2`
	expect := []int{0, 4, 4, 4}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8 5
7 4 2 4 8 2 1 4
5 4
1 10
3 2
8 11
7 7`
	expect := []int{5, 3, 6, 6, 9, 7}
	runSample(t, s, expect)
}
