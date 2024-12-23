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
	s := `1 2 3
10
15
5 5
5 15 20`
	expect := []int{1, 1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 4 5
8 100 8
10 150 20
2 32 1 1
9 200 51 50 10`
	expect := []int{1, 4, 2, 2, 1}
	runSample(t, s, expect)
}
