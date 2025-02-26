package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	reader := bufio.NewReader(strings.NewReader(expect))

	ans := readNNums(reader, len(res))

	if !reflect.DeepEqual(ans, res) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 7
2 4
5 1
2 3
3 4
4 1
5 3
3 5
`
	expect := "10 9 10 10 9 "

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 3
1 2
1 2
1 2
`
	expect := "5 6"

	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 1
3 2
`
	expect := "6 5 4 8 7 "

	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `50 20
4 18
39 33
49 32
7 32
38 1
46 11
8 1
3 31
30 47
24 16
33 5
5 21
3 48
13 23
49 50
18 47
40 32
9 23
19 39
25 12
`
	expect := "99 98 97 127 126 125 124 123 122 121 120 119 118 117 116 115 114 113 112 111 110 109 108 107 106 105 104 103 102 101 100 99 98 97 96 95 94 93 92 93 92 91 90 89 88 87 86 85 84 100  "

	runSample(t, s, expect)
}
