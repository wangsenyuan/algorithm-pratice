package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))

	n, m, a, b, res := process(reader)
	if len(res) == 0 {
		t.Fatalf("Sample result %v, not correct", res)
	}

	i1, j1, i2, j2 := res[0]-1, res[1]-1, res[2]-1, res[3]-1

	if i1 < 0 || i1+m > n || j1 < 0 || j1+m > m {
		t.Fatalf("Sample result %v, not correct", res)
	}

	if i2 < 0 || i2+m > m || j2 < 0 || j2+m > n {
		t.Fatalf("Sample result %v, not correct", res)
	}

	get := func(s []string, r int, c int) string {
		var buf bytes.Buffer
		for i := r; i < r+m; i++ {
			buf.WriteString(s[i][c : c+m])
		}
		return buf.String()
	}

	x := get(a, i1, j1)
	y := get(b, i2, j2)

	if x != y {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `10 5
somer
andom
noise
mayth
eforc
ebewi
thyou
hctwo
again
noise
somermayth
andomeforc
noiseebewi
againthyou
noisehctwo
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `2 2
ba
aa
ba
aa`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `3 2
ba
aa
cc
baa
acc`
	runSample(t, s)
}
