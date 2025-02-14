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
	s := `6 5 3
1 2 3`
	expect := []int{2, 3, 5}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 1 4
2 1 1 2`
	expect := []int{2, 2, 2, 2}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 3 1
3`
	expect := []int{2}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 2 4
2 1 1 1`
	expect := []int{2, 3, 3, 3}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `18 15 4
13 15 1 16`
	expect := []int{2, 4, 6, 8}
	runSample(t, s, expect)
}
