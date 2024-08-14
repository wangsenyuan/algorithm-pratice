package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	queries := make([][]int, k)
	for i := 0; i < k; i++ {
		queries[i] = readNNums(reader, 2)
	}

	res := solve(n, queries)
	var buf bytes.Buffer

	for _, x := range res {
		if x {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}
	fmt.Println(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(n int, queries [][]int) []bool {
	maze := make([]int, n+1)

	ans := make([]bool, len(queries))

	var cnt int

	check := func(c int) int {
		if maze[c]|maze[c+1] == 3 {
			return 1
		}
		return 0
	}

	process := func(r, c int) bool {
		cnt -= check(c)
		if c > 0 {
			cnt -= check(c - 1)
		}
		if c+1 < n {
			cnt -= check(c + 1)
		}
		if r == 0 {
			maze[c] ^= 1
		} else {
			maze[c] ^= 2
		}

		cnt += check(c)
		if c > 0 {
			cnt += check(c - 1)
		}
		if c+1 < n {
			cnt += check(c + 1)
		}

		return maze[0]&1 == 0 && maze[n-1]&2 == 0 && cnt == 0
	}

	for i, cur := range queries {
		ans[i] = process(cur[0]-1, cur[1]-1)
	}

	return ans
}

func solve1(n int, queries [][]int) []bool {
	// 只考虑第一行
	block := NewSegTree(n)

	maze := make([][]int, 2)
	for i := 0; i < 2; i++ {
		maze[i] = make([]int, n)
	}

	checkBlock := func(c int) int {
		if maze[0][c] == 0 {
			return 0
		}
		if c > 0 && maze[1][c-1] == 1 {
			return 1
		}
		if maze[1][c] == 1 {
			return 1
		}
		if c+1 < n && maze[1][c+1] == 1 {
			return 1
		}
		return 0
	}

	res := make([]bool, len(queries))

	process := func(r int, c int) bool {
		maze[r][c] ^= 1

		if r == 0 {
			block.Update(c, checkBlock(c))
		} else {
			// r == 1
			block.Update(c, checkBlock(c))
			if c > 0 {
				block.Update(c-1, checkBlock(c-1))
			}
			if c+1 < n {
				block.Update(c+1, checkBlock(c+1))
			}
		}

		if maze[0][0] == 1 || maze[1][n-1] == 1 {
			return false
		}
		return block.Query(0, n) == 0
	}

	for i, cur := range queries {
		res[i] = process(cur[0]-1, cur[1]-1)
	}

	return res
}

type SegTree []int

func NewSegTree(n int) SegTree {
	arr := make([]int, 2*n)
	return SegTree(arr)
}

func (tree SegTree) Update(p int, v int) {
	n := len(tree)
	n /= 2
	p += n
	tree[p] = v
	for p > 0 {
		tree[p>>1] = tree[p] + tree[p^1]
		p >>= 1
	}
}

func (tree SegTree) Query(l, r int) int {
	n := len(tree)
	n /= 2
	l += n
	r += n
	var res int
	for l < r {
		if l&1 == 1 {
			res += tree[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += tree[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}
