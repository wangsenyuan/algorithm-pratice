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
	s := `abacaba
5
2 1 4
1 4 b
1 5 b
2 4 6
2 1 7`
	expect := []int{3, 1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `dfcbbcfeeedbaea
15
1 6 e
1 4 b
2 6 14
1 7 b
1 12 c
2 6 8
2 1 6
1 7 c
1 2 f
1 10 a
2 7 9
1 10 a
1 14 b
1 1 f
2 1 11
`
	expect := []int{5, 2, 5, 2, 6}
	runSample(t, s, expect)
}
