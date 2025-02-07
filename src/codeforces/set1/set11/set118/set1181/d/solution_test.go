package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)
	n := len(res)
	reader = bufio.NewReader(strings.NewReader(expect))
	for i := range n {
		cur := readNum(reader)
		if cur != res[i] {
			t.Fatalf("Sample expect %s, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `6 4 10
3 1 1 1 2 2
7
8
9
10
11
12
13
14
15
16
`
	expect := `4
3
4
2
3
4
1
2
3
4
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 5 4
4 4 5 1
15
9
13
6
`
	expect := `5
3
3
3
`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `100 99 20
89 96 26 16 28 26 96 14 46 92 76 66 93 65 45 7 89 6 95 53 34 59 15 75 8 25 20 99 41 90 51 14 94 23 68 25 58 81 50 85 87 53 27 67 64 30 55 77 7 30 10 32 38 1 67 94 44 16 49 38 63 76 81 47 55 92 67 25 61 96 2 52 33 69 4 64 24 29 88 70 45 76 83 79 22 87 5 62 11 30 17 82 77 48 73 88 96 35 96 82
106
198
109
190
154
160
127
147
163
153
143
194
189
172
181
129
162
104
175
172
`
	expect := `19
86
36
74
32
39
97
22
42
31
18
80
73
52
63
1
41
13
57
52
`
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `20 9 20
9 9 9 9 9 9 9 9 9 9 9 9 9 9 9 9 9 9 9 9
172
27
41
96
137
111
33
36
67
173
176
150
72
75
38
81
186
180
169
40
`
	expect := `8
7
5
4
5
3
5
8
7
1
4
2
4
7
2
5
6
8
5
4
`
	runSample(t, s, expect)
}
