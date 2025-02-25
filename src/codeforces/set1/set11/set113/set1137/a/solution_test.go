package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	n := len(res)
	m := len(res[0])

	reader := bufio.NewReader(strings.NewReader(expect))
	for i := range n {
		cur := readNNums(reader, m)
		for j := 0; j < m; j++ {
			if res[i][j] != cur[j] {
				t.Fatalf("Sample expect %s, but got %v", expect, res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `2 3
1 2 1
2 1 2
`
	expect := `2 2 2 
2 2 2 
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2
1 2
3 4
`
	expect := `2 3 
3 2 
`
	runSample(t, s, expect)
}
