package p1079

import (
	"bytes"
	"sort"
)

var C [][]int

func init() {
	C = make([][]int, 8)

	for i := 0; i < 8; i++ {
		C[i] = make([]int, 8)
	}

	C[0][0] = 1

	for i := 1; i < 8; i++ {
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}
}

func numTilePossibilities(tiles string) int {
	n := len(tiles)

	bs := make([]int, n)

	for i := 0; i < n; i++ {
		bs[i] = int(tiles[i] - 'A')
	}

	sort.Ints(bs)

	mem := make(map[string]bool)

	cnt := make([]int, 26)
	N := 1 << uint(n)
	var buf bytes.Buffer

	var res int

	for mask := 1; mask < N; mask++ {
		buf.Reset()

		for i := 0; i < 26; i++ {
			cnt[i] = 0
		}

		for i := 0; i < n; i++ {
			if mask&(1<<uint(i)) > 0 {
				buf.WriteByte(byte('A' + bs[i]))
				cnt[bs[i]]++
			}
		}
		s := buf.String()

		if mem[s] {
			continue
		}
		mem[s] = true

		m := len(s)

		num := 1
		for i := 0; i < 26; i++ {
			num *= C[m][cnt[i]]
			m -= cnt[i]
		}

		res += num
	}

	return res
}

func numTilePossibilities1(tiles string) int {
	n := len(tiles)
	mem := make(map[string]bool)

	var loop func(s string, mask int, r []byte, n int)

	loop = func(s string, mask int, r []byte, n int) {
		if mask == 1<<uint(len(s))-1 {
			if n > 0 {
				mem[string(r[:n])] = true
			}
			return
		}

		for i := 0; i < len(s); i++ {
			if mask&(1<<uint(i)) == 0 {
				next := mask | (1 << uint(i))
				loop(s, next, r, n)
				r[n] = s[i]
				loop(s, next, r, n+1)
			}
		}
	}

	N := 1 << uint(n)

	r := make([]byte, n)
	var buf bytes.Buffer

	for i := 1; i < N; i++ {
		buf.Reset()
		for j := 0; j < n; j++ {
			if i&(1<<uint(j)) > 0 {
				buf.WriteByte(tiles[j])
			}
		}

		s := buf.String()

		loop(s, 0, r, 0)
	}

	return len(mem)
}
