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
	s := `6 3
3 4 6`
	expect := []int{6, 4, 4, 3, 4, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 1
2`
	runSample(t, s, nil)
}
