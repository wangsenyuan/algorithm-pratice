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
	s := `5
1 2 5 905 2000000
15
- 2
? 2
? 1
- 1
? 1
+ 4
+ 2
? 2
+ 6
- 4
+ 7
? 2
? 3
? 4
? 2000000`
	expect := []int{2, 2, 1, 6, 3, 8, 8, 2000001}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
3 4 5 6 8
9
? 5
- 5
? 5
+ 1
? 2
- 6
- 8
+ 6
? 5`
	expect := []int{9, 9, 9, 7}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
6 7 8 9 10
10
? 5
- 6
? 4
- 10
+ 5
- 8
+ 3
+ 2
- 3
+ 10`
	expect := []int{1, 1}
	runSample(t, s, expect)
}
