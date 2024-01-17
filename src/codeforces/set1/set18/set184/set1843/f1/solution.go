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
		queries := make([]string, n)
		for i := 0; i < n; i++ {
			queries[i] = readString(reader)
		}
		res := solve(queries)
		for _, x := range res {
			if x {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
		}
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

func solve(queries []string) []bool {
	//n := len(queries)

	type node [4]int

	var nodes []node
	// max suf, max ans, min suf, min ans
	nodes = append(nodes, [...]int{1, 1, 0, 0})
	id := 1

	add := func(v int, x int) {
		pref := nodes[v]
		cur := [...]int{0, 0, 0, 0}
		cur[2] = min(0, pref[2]+x)
		cur[0] = max(0, pref[0]+x)
		cur[3] = min(pref[3], cur[2])
		cur[1] = max(pref[1], cur[0])
		nodes = append(nodes, cur)
		id++
	}

	query := func(u int, v int, k int) bool {
		// u = 0
		cur := nodes[v]

		return cur[3] <= k && k <= cur[1]
	}

	var res []bool

	for _, cur := range queries {
		s := []byte(cur)
		if cur[0] == '+' {
			var v, x int
			pos := readInt(s, 2, &v)
			readInt(s, pos+1, &x)
			add(v-1, x)
		} else {
			var u, v, k int
			pos := readInt(s, 2, &u)
			pos = readInt(s, pos+1, &v)
			readInt(s, pos+1, &k)
			res = append(res, query(u-1, v-1, k))
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
