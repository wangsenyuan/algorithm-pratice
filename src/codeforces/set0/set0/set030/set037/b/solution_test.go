package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	hp, reg, scrolls, time, res := process(bufio.NewReader(strings.NewReader(s)))

	if time != expect {
		t.Fatalf("Sample expect %d, but got %d %v", expect, time, res)
	}
	if expect < 0 {
		return
	}
	var sum, prev int
	cur := hp
	for _, op := range res {
		x, i := op[0], op[1]-1
		cur = cur - (x-prev)*sum
		cur = min(hp, cur+reg*(x-prev))
		if cur <= 0 {
			t.Fatalf("Sample result %v not correct, the boss already down", res)
		}
		pow, dmg := scrolls[i][0], scrolls[i][1]
		if pow*hp < cur*100 {
			t.Fatalf("Sample result %v not correct, it can't use scroll %d, but it used", res, i)
		}
		sum += dmg
		prev = x
	}
	cur -= (expect - prev) * sum
	cur += (expect - prev) * reg
	if cur > 0 {
		t.Fatalf("Sample result %v, not correct, the boss is still alive", res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 10 3
100 3
99 1
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 100 10
100 11
90 9
`
	expect := 19
	runSample(t, s, expect)
}
