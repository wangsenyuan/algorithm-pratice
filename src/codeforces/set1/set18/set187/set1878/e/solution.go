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
		n := readNum(reader)
		a := readNNums(reader, n)
		m := readNum(reader)
		queries := make([][]int, m)
		for i := 0; i < m; i++ {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(a, queries)

		s := fmt.Sprintf("%v", res)
		s = s[1 : len(s)-1]

		buf.WriteString(fmt.Sprintf("%s\n", s))
	}

	fmt.Print(buf.String())
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

const H = 31

func solve(a []int, queries [][]int) []int {
	n := len(a)
	next := make([][]int, n+1)
	for i := range n + 1 {
		next[i] = make([]int, H)
		for j := 0; j < H; j++ {
			next[i][j] = n
		}
	}

	for i := n - 1; i >= 0; i-- {
		copy(next[i], next[i+1])
		for j := 0; j < H; j++ {
			if (a[i]>>j)&1 == 0 {
				next[i][j] = i
			}
		}
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		l, k := cur[0]-1, cur[1]
		num := a[l]
		if num < k {
			ans[i] = -1
			continue
		}
		ans[i] = l + 1
		r := n
		for j := H - 1; j >= 0; j-- {
			x := (k >> j) & 1
			if x == 0 {
				ans[i] = max(ans[i], min(r, next[l][j]))
			} else {
				r = min(r, next[l][j])
			}
		}
		ans[i] = max(ans[i], r)
	}
	return ans
}

func solve2(a []int, queries [][]int) []int {
	n := len(a)
	pt := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		pt[i] = make([]int, H)
	}
	for j := 0; j < H; j++ {
		pt[n][j] = n
	}

	for i := n - 1; i >= 0; i-- {
		copy(pt[i], pt[i+1])
		for j := H - 1; j >= 0; j-- {
			if (a[i]>>j)&1 == 0 {
				pt[i][j] = i
			}
		}
	}

	get := func(l int, r int) int {
		var res int
		for j := H - 1; j >= 0; j-- {
			if pt[l][j] > r {
				res |= 1 << j
			}
		}
		return res
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		l, k := cur[0]-1, cur[1]
		num := a[l]
		if num < k {
			ans[i] = -1
			continue
		}
		ans[i] = search(l, n, func(r int) bool {
			return get(l, r) < k
		})
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func solve1(a []int, queries [][]int) []int {
	// binary lift
	n := len(a)

	arr := make([]int, 2*n)
	copy(arr[n:], a)

	for i := n - 1; i > 0; i-- {
		arr[i] = arr[i*2] & arr[i*2+1]
	}

	get := func(l int, r int) int {
		l += n
		r += n
		res := (1 << H) - 1

		for l < r {
			if l&1 == 1 {
				res &= arr[l]
				l++
			}
			if r&1 == 1 {
				r--
				res &= arr[r]
			}
			l >>= 1
			r >>= 1
		}

		return res
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		l, k := cur[0]-1, cur[1]
		if a[l] < k {
			ans[i] = -1
			continue
		}
		r := search(l, n, func(r int) bool {
			return get(l, r+1) < k
		})
		ans[i] = r
	}

	return ans
}

func search(l int, r int, f func(int) bool) int {
	for l < r {
		mid := (l + r) >> 1
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
