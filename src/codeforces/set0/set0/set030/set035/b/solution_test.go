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
	s := `2 2 9
+1 1 1 cola
+1 1 1 fanta
+1 1 1 sevenup
+1 1 1 whitekey
-1 cola
-1 fanta
-1 sevenup
-1 whitekey
-1 cola
`
	expect := [][]int{
		{1, 1},
		{1, 2},
		{2, 1},
		{2, 2},
		{-1, -1},
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2 8
+1 1 1 cola
-1 cola
+1 1 1 fanta
-1 fanta
+1 1 1 sevenup
-1 sevenup
+1 1 1 whitekey
-1 whitekey
`
	expect := [][]int{
		{1, 1},
		{1, 1},
		{1, 1},
		{1, 1},
	}
	runSample(t, s, expect)
}