package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	reader = bufio.NewReader(strings.NewReader(expect))
	for _, cur := range res {
		s := readString(reader)
		if s == "Impossible" && len(cur) != 0 {
			t.Fatalf("Sample expect %s, but got %v", s, cur)
		}
		if s != "Impossible" && len(cur) == 0 {
			t.Fatalf("Sample expect %s, but got %v", s, cur)
		}
		if s == "Impossible" {
			continue
		}
		x := fmt.Sprintf("%d %d", cur[0], cur[1])
		if x != s {
			t.Fatalf("Sample expect %s, but got %v", s, cur)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5
30 50 70 20 60
NYYNN
NNYNN
NNNYY
YNNNN
YNNNN
3
1 3
3 1
4 5`
	expect := `1 100
2 160
3 180`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
100 100
NN
NN
1
1 2`
	expect := `Impossible`
	runSample(t, s, expect)
}
