package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))

	p, q := process(reader)

	type trip struct {
		first  int
		second int
		third  int
	}

	convert := func(arr []int) trip {
		a, b, c := arr[0], arr[1], arr[2]
		sum := a + b + c
		a, c = min(a, b, c), max(a, b, c)
		b = sum - a - c
		return trip{a, b, c}
	}
	freq := make(map[trip]int)
	for i := 0; i+2 < len(p); i++ {
		freq[convert(p[i:i+3])]++
	}

	for _, cur := range q {
		freq[convert(cur)]--
	}

	for _, v := range freq {
		if v != 0 {
			t.Fatalf("Sample result %v, not correct", p)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5
4 3 2
2 3 5
4 1 2`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `5
4 5 3
2 5 1
1 3 5`
	runSample(t, s)
}
