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
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `7
3 4 1 2 2 1 1
`
	expect := []int{3, 8, 2, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 1 3 1 1
`
	expect := []int{3, 4}
	runSample(t, s, expect)
}


func TestSample3(t *testing.T) {
	s := `5
10 40 20 50 30
`
	expect := []int{10, 40, 20, 50, 30}
	runSample(t, s, expect)
}
