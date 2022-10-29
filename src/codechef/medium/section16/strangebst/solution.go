package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	W := readNNums(reader, n)
	S := readNNums(reader, n)
	res := solve(n, W, S)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(n int, W []int, S []int) int64 {
	type Node struct {
		support int
		weight  int
	}

	nodes := make([]Node, n)

	for i := 0; i < n; i++ {
		nodes[i] = Node{S[i], W[i]}
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].weight < nodes[j].weight
	})

	type Pair struct {
		first  int
		second int
	}

	min_pair := func(a, b Pair) Pair {
		if a.first < b.first {
			return a
		}
		return b
	}

	arr := make([]Pair, 2*n)

	for i := n; i < 2*n; i++ {
		arr[i] = Pair{nodes[i-n].support - nodes[i-n].weight, i - n}
	}

	for i := n - 1; i > 0; i-- {
		arr[i] = min_pair(arr[2*i], arr[2*i+1])
	}

	get := func(l, r int) Pair {
		l += n
		r += n
		res := Pair{1 << 30, -1}
		for l < r {
			if l&1 == 1 {
				res = min_pair(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = min_pair(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	pref := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		pref[i] = pref[i-1] + int64(nodes[i-1].weight)
	}

	var res int64

	var dfs func(l int, r int)

	dfs = func(l int, r int) {
		x := get(l, r).second
		kx := pref[r] - pref[l] + int64(nodes[x].support-nodes[x].weight)
		res = max2(res, kx)
		if l < x {
			dfs(l, x)
		}
		if x+1 < r {
			dfs(x+1, r)
		}
	}

	dfs(0, n)

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max2(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
