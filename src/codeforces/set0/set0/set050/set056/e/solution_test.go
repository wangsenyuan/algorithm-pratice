package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	result := process(reader)

	// Parse expected result string into an array of integers
	expectReader := bufio.NewReader(strings.NewReader(expect))
	expectArr := readNNums(expectReader, len(result))

	// Compare result with expected using reflect.DeepEqual
	if !reflect.DeepEqual(result, expectArr) {
		t.Fatalf("Expected %v, but got %v", expectArr, result)
	}
}

func TestSample1(t *testing.T) {
	s := `4
16 5
20 5
10 10
18 2
`
	expect := `3 1 4 1`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
0 10
1 5
9 10
15 10
`
	expect := `4 1 2 1`
	runSample(t, s, expect)
}
