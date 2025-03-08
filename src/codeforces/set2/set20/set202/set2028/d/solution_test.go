package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	ok, label, nums, n, q, k, j := process(bufio.NewReader(strings.NewReader(s)))

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if !ok {
		return
	}
	cur := 1
	for i := range label {
		var a []int
		if label[i] == 'q' {
			a = q
		} else if label[i] == 'k' {
			a = k
		} else {
			a = j
		}
		v := nums[i]
		if v < cur || a[cur-1] < a[v-1] {
			t.Fatalf("can't exchange for %d %d with %c", cur, v, label[i])
		}
		cur = v
	}

	if cur != n {
		t.Fatalf("Sample result not correct")
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 3 2
2 1 3
1 2 3`
	expect := true
	runSample(t, s, expect)
}
func TestSample2(t *testing.T) {
	s := `4
2 3 1 4
1 2 3 4
1 4 2 3`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
1 2 3
1 2 3
3 1 2`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3
1 2 3
1 2 3
2 3 1`
	expect := true
	runSample(t, s, expect)
}
