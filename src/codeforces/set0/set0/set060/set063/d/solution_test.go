package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, island, p := process(reader)
	// n := len(p)
	a, b, c, d := island[0], island[1], island[2], island[3]

	if b < d {
		for i := b; i < d; i++ {
			for j := 0; j < a; j++ {
				if res[i][j] != '.' {
					t.Fatalf("Sample result %v, not in the right shape", res)
				}
			}
		}
	} else if b > d {
		for i := d; i < b; i++ {
			for j := a; j < a+c; j++ {
				if res[i][j] != '.' {
					t.Fatalf("Sample result %v, not in the right shape", res)
				}
			}
		}
	}

	que := make([]int, len(res)*len(res[0]))
	n := len(res)
	m := len(res[0])

	vis := make([][]bool, n)
	for i := range len(res) {
		vis[i] = make([]bool, m)
	}

	var dd = []int{-1, 0, 1, 0, -1}

	bfs := func(x int, y int) int {
		var head, tail int
		que[head] = x*m + y
		head++
		vis[x][y] = true
		for tail < head {
			u, v := que[tail]/m, que[tail]%m
			tail++
			for i := 0; i < 4; i++ {
				r, c := u+dd[i], v+dd[i+1]
				if r >= 0 && r < n && c >= 0 && c < m && res[r][c] == res[x][y] && !vis[r][c] {
					que[head] = r*m + c
					head++
					vis[r][c] = true
				}
			}
		}

		return head
	}

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			if res[i][j] == '.' || vis[i][j] {
				continue
			}
			id := int(res[i][j] - 'a')
			cnt := bfs(i, j)
			if cnt != p[id] {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}

}

func TestSample1(t *testing.T) {
	s := `3 4 2 2 3
5 8 3
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 2 1 4 4
1 2 3 4
`
	runSample(t, s)
}
