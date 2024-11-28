package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)

	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := readNNums(reader, n)
	m := readNum(reader)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = readNNums(reader, 2)
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	n := len(a)
	pref := make([]int, n+1)

	cnt := make([]int, n+1)
	for i, num := range a {
		pref[i+1] = pref[i] + num
		cnt[i+1] = cnt[i] + n - i
	}

	locate := func(pos int) (r int, c int) {
		// sum[r] >=¸ pos
		r = sort.SearchInts(cnt, pos)
		if cnt[r] > pos {
			r--
		}
		c = pos - cnt[r]
		c += r
		return
	}

	psum := make([]int, n+1)

	for i := 1; i <= n; i++ {
		psum[i] = psum[i-1] + pref[i]
	}

	get := func(r int, c int) int {
		// 计算在r行，从c开始的前缀和
		// sum[2.2] + sum[2.3] + .. sum[2.c-1] + sum[2..c] + ... sum[2...n-1]
		res := psum[c+1] - psum[r]
		res -= (c - r + 1) * pref[r]
		return res
	}

	pref2 := make([]int, n+1)

	for i := range n {
		pref2[i+1] = pref2[i] + get(i, n-1)
	}

	find := func(l int, r int) int {
		x, y := locate(l)
		u, v := locate(r)
		var res int
		if x == u {
			// same row
			res = get(x, v) - get(x, y-1)
		} else {
			// x < u
			res = pref2[u] - pref2[x+1]
			res += get(x, n-1) - get(x, y-1)
			res += get(u, v)
		}

		return res
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		ans[i] = find(cur[0]-1, cur[1]-1)
	}

	return ans
}
