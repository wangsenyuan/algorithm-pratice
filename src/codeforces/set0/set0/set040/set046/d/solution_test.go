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
	s := `30 1 2
6
1 5
1 4
1 5
2 2
1 5
1 4
`
	expect := []int{0, 6, 11, 17, 23}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `30 1 1
6
1 5
1 4
1 5
2 2
1 5
1 4
`
	expect := []int{0, 6, 11, 17, 6}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 1 1
1
1 12
`
	expect := []int{-1}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 1 1
1
1 10
`
	expect := []int{0}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `10 1 1
2
1 3
1 6
`
	expect := []int{0, 4}
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `10 1 1
5
1 4
2 1
1 3
2 3
1 1
`
	expect := []int{0, 0, 0}
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `50 2 3
15
1 7
1 6
1 1
2 3
2 1
2 2
1 1
1 4
1 6
1 2
1 8
1 6
2 7
1 8
2 9
`
	expect := []int{0, 9, 17, 0, 3, 9, 17, 21, 31, 39}
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `1000 1 100
84
1 33
1 13
2 1
2 2
1 6
1 18
1 6
1 6
1 17
2 7
1 47
1 9
1 17
1 15
1 33
2 15
2 12
2 8
1 74
2 14
2 5
1 64
2 22
1 45
2 19
1 95
2 24
1 29
1 33
1 15
1 66
1 30
2 9
2 13
2 29
1 28
1 9
2 26
1 70
2 31
1 31
2 6
1 60
1 13
1 93
1 31
1 62
1 77
1 97
1 41
2 37
1 30
1 97
2 45
1 96
1 71
2 41
2 43
1 24
1 78
1 48
1 22
2 47
1 17
1 71
1 39
1 11
1 31
2 44
1 19
1 81
2 70
2 59
1 23
2 32
1 56
2 46
1 17
1 41
1 31
1 90
1 42
2 36
1 47
`
	expect := []int{0, 34, 0, 7, 26, 33, 40, 58, 106, 116, 134, 150, 150, 225, 225, 271, 134, 367, 401, 417, 484, 515, 544, 164, 235, 554, 267, 615, 709, 741, 804, 882, -1, -1, -1, -1, -1, 544, -1, -1, 569, 281, -1, -1, 592, -1, 980, -1, -1, -1, 604, 622, 664, -1, -1, 417}
	runSample(t, s, expect)
}
