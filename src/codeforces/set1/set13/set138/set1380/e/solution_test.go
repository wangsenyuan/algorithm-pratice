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
		t.Fatalf("Sample %s expect %v, but got %v", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `7 4
1 2 3 3 1 4 3
3 1
2 3
2 4
`
	expect := []int{5, 4, 2, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 10
2 10 8 7 9 3 4 6 5 1
10 4
10 1
10 6
10 8
10 2
10 9
10 7
10 3
10 5
`
	expect := []int{9, 9, 9, 8, 7, 6, 6, 4, 2, 0}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 5
3 1 5 3 1 5 1 2 4 2
4 5
4 3
4 2
4 1
`
	expect := []int{9, 9, 8, 6, 0}
	runSample(t, s, expect)
}
func TestSample5(t *testing.T) {
	s := `10 4
1 4 3 4 2 2 4 3 4 2 
2 4
1 3
2 1`
	expect := []int{8, 5, 5, 0}
	runSample(t, s, expect)
}
