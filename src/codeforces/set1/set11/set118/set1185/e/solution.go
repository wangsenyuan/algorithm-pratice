package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		ok, res, _ := process(reader)
		if !ok {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString(fmt.Sprintf("YES\n%d\n", len(res)))
			for i := 0; i < len(res); i++ {
				for j := 0; j < len(res[i]); j++ {
					buf.WriteString(fmt.Sprintf("%d ", res[i][j]))
				}
				buf.WriteByte('\n')
			}
		}
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
}

func process(reader *bufio.Reader) (ok bool, res [][]int, grid []string) {
	n, m := readTwoNums(reader)
	grid = make([]string, n)
	for i := range n {
		grid[i] = readString(reader)[:m]
	}
	ok, res = solve(grid)
	return
}

const inf = 1 << 30

func solve(grid []string) (bool, [][]int) {
	n := len(grid)
	m := len(grid[0])
	rows := make([]*SegTree, n)
	for i := range n {
		rows[i] = NewSegTree(m)
	}
	cols := make([]*SegTree, m)
	for i := range m {
		cols[i] = NewSegTree(n)
	}

	pos := make([][]int, 26)
	for i := range 26 {
		pos[i] = []int{inf, inf, -1, -1}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '.' {
				rows[i].Update(j, -1)
				cols[j].Update(i, -1)
			} else {
				x := int(grid[i][j] - 'a')
				rows[i].Update(j, x)
				cols[j].Update(i, x)
				pos[x][0] = min(pos[x][0], i)
				pos[x][1] = min(pos[x][1], j)
				pos[x][2] = max(pos[x][2], i)
				pos[x][3] = max(pos[x][3], j)
			}
		}
	}
	var res [][]int
	var prev []int
	for x := 25; x >= 0; x-- {
		if pos[x][0] == inf {
			if len(prev) > 0 {
				res = append(res, prev)
			}
			continue
		}
		a, b, c, d := pos[x][0], pos[x][1], pos[x][2], pos[x][3]
		if a < c && b < d {
			return false, nil
		}

		if a == c {
			// 同一行
			y := rows[a].Get(b, d+1)
			if y != x {
				return false, nil
			}
		} else if b == d {
			// 同一列
			y := cols[b].Get(a, c+1)
			if y != x {
				return false, nil
			}
		} else {
			return false, nil
		}
		prev = []int{a + 1, b + 1, c + 1, d + 1}
		res = append(res, prev)
	}
	reverse(res)
	return true, res
}

func reverse(arr [][]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type SegTree struct {
	size int
	arr  []int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = inf
	}
	return &SegTree{n, arr}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = min(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := inf
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = min(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
