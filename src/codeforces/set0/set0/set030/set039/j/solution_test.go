package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `abdrakadabra
abrakadabra
`
	expect := []int{3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `aa
a
`
	expect := []int{1, 2}
	runSample(t, s, expect)
}
