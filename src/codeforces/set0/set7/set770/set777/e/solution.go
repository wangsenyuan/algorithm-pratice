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
	rings := make([][]int, n)
	for i := 0; i < n; i++ {
		rings[i] = readNNums(reader, 3)
	}

	res := solve(rings)

	fmt.Println(res)
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

func solve(rings [][]int) int {
	n := len(rings)
	arr := make([]int, n*2)
	for i := 0; i < n; i++ {
		arr[i*2] = rings[i][0]
		arr[i*2+1] = rings[i][1]
	}
	id := compress(arr)

	m := len(id)

	tr := make(fenwick, m+10)

	sort.Slice(rings, func(i, j int) bool {
		x := rings[i]
		y := rings[j]
		return x[1] > y[1] || x[1] == y[1] && x[0] > y[0]
	})

	var best int

	for _, cur := range rings {
		a, b, h := cur[0], cur[1], cur[2]
		a = id[a]
		b = id[b]
		// 前面的都可以用来放置当前的
		// 找到前面的比b小的内圈的最好的答案
		tmp := tr.pre(b - 1)
		tmp += h
		best = max(best, tmp)
		tr.update(a, tmp)
	}

	return best
}

func compress(a []int) map[int]int {
	pos := make(map[int]int)
	sort.Ints(a)
	for _, x := range a {
		if _, ok := pos[x]; ok {
			continue
		}
		pos[x] = len(pos)
	}
	return pos
}

type fenwick []int

func (f fenwick) update(i int, val int) {
	for i++; i < len(f); i += i & -i {
		f[i] = max(f[i], val)
	}
}
func (f fenwick) pre(i int) (res int) {
	for i++; i > 0; i &= i - 1 {
		res = max(res, f[i])
	}
	return
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
