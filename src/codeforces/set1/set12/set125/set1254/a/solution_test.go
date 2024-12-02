package main

import (
	"bufio"
	"strings"
	"testing"
)

func get(id byte) int {
	return strings.IndexByte(letters, id)
}

func runSample(t *testing.T, s string, expect []string) {
	reader := bufio.NewReader(strings.NewReader(s))

	res, land, _ := process(reader)

	check := func(res []string) int {
		buf := make([][]byte, len(res))
		for i := range len(res) {
			buf[i] = []byte(res[i])
		}

		n := len(res)
		m := len(res[0])
		cnt := make(map[int]int)

		que := make([]int, n*m)

		bfs := func(r int, c int, x byte) {
			buf[r][c] = '0'
			var head, tail int
			que[head] = r*m + c
			head++
			for tail < head {
				u, v := que[tail]/m, que[tail]%m
				if land[u][v] == 'R' {
					cnt[get(x)]++
				}
				tail++
				for i := 0; i < 4; i++ {
					nu, nv := u+dd[i], v+dd[i+1]
					if nu >= 0 && nu < n && nv >= 0 && nv < m && buf[nu][nv] == x {
						buf[nu][nv] = '0'
						que[head] = nu*m + nv
						head++
					}
				}
			}
		}

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if buf[i][j] != '0' {
					x := buf[i][j]
					if cnt[get(x)] > 0 {
						t.Fatalf("Sample result %v, not correct, it has dis-connected regions", res)
					}
					bfs(i, j, x)
				}
			}
		}

		mn, mx := n, 0
		for _, v := range cnt {
			mn = min(mn, v)
			mx = max(mx, v)
		}

		return mx - mn
	}

	x := check(expect)
	y := check(res)

	if x != y {
		t.Fatalf("Sample expect %v, but got %v with diff %d - %d", expect, res, x, y)
	}
}

func TestSample1(t *testing.T) {
	s := `3 5 3
..R..
...R.
....R`
	expect := []string{
		"11122",
		"22223",
		"33333",
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 4 6
R..R
R..R
RRRR
RRRR
R..R
R..R`
	expect := []string{
		"aacc",
		"aBBc",
		"aBBc",
		"CbbA",
		"CbbA",
		"CCAA",
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 5 4
RRR..
R.R..
RRR..
R..R.
R...R`
	expect := []string{
		"11114",
		"22244",
		"32444",
		"33344",
		"33334",
	}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 31 62
RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR
RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR`
	expect := []string{
		"abcdefghijklmnopqrstuvwxyzABCDE",
		"FGHIJKLMNOPQRSTUVWXYZ0123456789",
	}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `2 6 3
RRRRRR
RRRRRR`
	expect := []string{
		"000011",
		"222211",
	}
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `1 5 3
RRRRR`
	expect := []string{
		"EEWWQ",
	}
	runSample(t, s, expect)
}
