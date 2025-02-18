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
	s := `4
2 3 1 4`
	expect := []int{3, 3, 3, 4}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
5 4 3 2 1`
	expect := []int{5, 5, 5, 5, 5}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
2 1 1 3`
	expect := []int{2, 2, 2, 3}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4
1 1 3 1`
	expect := []int{1, 1, 3, 3}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `8
2 4 1 6 3 8 5 7`
	// 这个例子表明还是有点复杂的
	expect := []int{8, 8, 8, 8, 8, 8, 8, 8}
	runSample(t, s, expect)
}
