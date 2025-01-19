package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	n := len(res)

	reader = bufio.NewReader(strings.NewReader(expect))

	arr := readNNums(reader, n)

	if !reflect.DeepEqual(res, arr) {
		t.Fatalf("Sample expect %v, but got %v", arr, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4
1 1 2
3 2 1
1 3
2 3
1 1
2 1`
	expect := "2 3 3 6 6"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 8
1 4 2 7 3 5
7 6 5 6 3 3
2 5
1 6
1 5
1 5
1 5
2 3
2 3
1 6`
	expect := "840 840 1008 1344 1680 2016 2016 2016 2352"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `13 8
7 7 6 6 5 5 5 2 2 3 4 5 1
1 4 1 9 6 6 9 1 5 1 3 8 4
2 2
2 11
2 4
2 4
1 7
1 1
2 12
1 5`
	expect := "2116800 2646000 3528000 3528000 3528000 4233600 4838400 4838400 4838400"
	runSample(t, s, expect)
}
