package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	ok, res, a := process(reader)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t %v", expect, ok, res)
	}

	if !expect {
		return
	}

	n := len(a)
	m := len(a[0])
	cnt := make([][]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = make([]int, m)
	}

	for _, cur := range res {
		x, y, r := cur[0]-1, cur[1]-1, cur[2]
		for i := 1; i <= r; i++ {
			cnt[x-i][y]++
			cnt[x+i][y]++
			cnt[x][y-i]++
			cnt[x][y+i]++
		}
		cnt[x][y]++
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == '*' && cnt[i][j] == 0 {
				t.Fatalf("Sample result %v, not correct", res)
			}
			if a[i][j] == '.' && cnt[i][j] > 0 {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `6 8
....*...
...**...
..*****.
...**...
....*...
........
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 5
.*...
****.
.****
..**.
.....
`
	expect := true
	runSample(t, s, expect)
}


func TestSample3(t *testing.T) {
	s := `5 5
.*...
***..
.*...
.*...
.....
`
	expect := false
	runSample(t, s, expect)
}