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
	s := `4 4 2
2 3
1 2
1 3
1 4
`
	expect := []int{0, 0, 3, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 8 2
2 1
4 2
5 4
5 2
4 3
5 1
4 1
3 2
`
	expect := []int{0, 0, 0, 3, 3, 4, 4, 5}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 7 2
1 5
3 2
2 5
3 4
1 2
5 3
1 3
`
	expect := []int{0, 0, 0, 0, 3, 4, 4}
	runSample(t, s, expect)
}
