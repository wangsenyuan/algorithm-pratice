package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	region, x, res := process(reader)

	if x != expect {
		t.Fatalf("Sample expect %d, but got %d, %v", expect, x, res)
	}

	if x == 0 {
		if !reflect.DeepEqual(region, res) {
			t.Fatalf("Sample result %v, not correct", res)
		}
		return
	}
	n := len(region)
	m := len(region[0])

	cnt := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		cnt[i] = make([]int, m+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if res[i][j] == 'X' {
				cnt[i-x][j-x]++
				cnt[i-x][j+x+1]--
				cnt[i+x+1][j-x]--
				cnt[i+x+1][j+x+1]++
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i > 0 {
				cnt[i][j] += cnt[i-1][j]
			}
			if j > 0 {
				cnt[i][j] += cnt[i][j-1]
			}
			if i > 0 && j > 0 {
				cnt[i][j] -= cnt[i-1][j-1]
			}
			if region[i][j] == 'X' != (cnt[i][j] > 0) {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 6
XXXXXX
XXXXXX
XXXXXX
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 10
.XXXXXX...
.XXXXXX...
.XXXXXX...
.XXXXXX...
.XXXXXXXX.
...XXXXXX.
...XXXXXX.
...XXXXXX.
...XXXXXX.
..........
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 5
X....
..XXX
..XXX
..XXX
`
	expect := 0
	runSample(t, s, expect)
}
