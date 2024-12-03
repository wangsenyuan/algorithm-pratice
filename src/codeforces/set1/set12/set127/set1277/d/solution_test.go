package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	k, res, words := process(reader)

	if k != expect {
		t.Fatalf("Sample expect %d, but got %d %v %v", expect, k, res, words)
	}

	if expect < 0 {
		return
	}

	for _, i := range res {
		words[i-1] = reverse(words[i-1])
	}
	cnt := make([]int, 2)
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
		if freq[w] > 1 {
			t.Fatalf("Sample result %v, not correct", res)
		}
		x := int(w[0] - '0')
		y := int(w[len(w)-1] - '0')
		if x == y {
			continue
		}
		cnt[x]++
	}

	if abs(cnt[0]-cnt[1]) > 1 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
0001
1000
0011
0111`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
010
101
0`
	expect := -1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
00000
00001`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4
01
001
0001
00001`
	expect := 2
	runSample(t, s, expect)
}
