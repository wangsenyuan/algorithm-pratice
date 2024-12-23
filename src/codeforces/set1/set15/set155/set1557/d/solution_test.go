package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, n, segs := process(reader)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	rows := make([][]pair, n+1)
	for _, seg := range segs {
		i, l, r := seg[0], seg[1], seg[2]
		rows[i] = append(rows[i], pair{l, r})
	}

	for i := 1; i <= n; i++ {
		slices.SortFunc(rows[i], func(a, b pair) int {
			return a.first - b.first
		})
	}

	var prev int
	var j int
	for i := 1; i <= n; i++ {
		if j < len(res) && res[j] == i {
			// 这行被删除了
			j++
			continue
		}
		// 这行被保留的
		if prev > 0 {
			ok := false
			for u, v := 0, 0; u < len(rows[prev]) && v < len(rows[i]); {
				a, b := rows[prev][u].first, rows[prev][u].second
				c, d := rows[i][v].first, rows[i][v].second
				if max(a, c) <= min(b, d) {
					ok = true
					break
				}
				if b < d {
					u++
				} else {
					v++
				}
			}
			if !ok {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
		prev = i
	}
}

func TestSample1(t *testing.T) {
	s := `3 6
1 1 1
1 7 8
2 7 7
2 15 15
3 1 1
3 15 15
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 4
1 2 3
2 4 6
3 3 5
5 1 1
`
	expect := 3
	runSample(t, s, expect)
}
