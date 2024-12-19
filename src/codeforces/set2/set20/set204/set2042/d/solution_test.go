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
	s := `3
3 8
2 5
4 5`
	expect := []int{0, 0, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
42 42
1 1000000000`
	expect := []int{999999999, 0}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
42 42
1 1000000000
42 42`
	expect := []int{0, 0, 0}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6
1 10
3 10
3 7
5 7
4 4
1 2`
	expect := []int{0, 2, 3, 2, 4, 8}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `2
1 1
1 8`
	expect := []int{7, 0}
	runSample(t, s, expect)
}
