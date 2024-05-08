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

	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}

	res := solve(n, edges)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}

	var buf bytes.Buffer

	buf.WriteString("YES\n")

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d %d\n", res[i][0], res[i][1]))
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

func solve(n int, edges [][]int) [][]int {
	nodes := make([]int, n-1)
	cnt := make([]int, n+1)
	for i, edge := range edges {
		u, v := edge[0], edge[1]
		u, v = min(u, v), max(u, v)

		if v != n {
			return nil
		}
		nodes[i] = u
		cnt[u]++
	}

	sort.Ints(nodes)
	var cur int
	for i := 1; i < n; i++ {
		cur += cnt[i]
		if cur > i {
			return nil
		}
	}

	for i, j := 1, 0; i < n; i++ {

		if j < len(nodes) {
			x := nodes[j]
			for j < len(nodes) && nodes[j] == x {
				j++
			}
			if x == i {
				j++
				continue
			}
		}
	}

	prev := -1
	var res [][]int
	p := 1
	vis := make([]bool, n+1)
	for i := 0; i < n-1; {
		j := i
		for i < n-1 && nodes[i] == nodes[j] {
			i++
		}
		tmp := i - j
		if tmp > 0 {
			vis[nodes[j]] = true
			if prev != -1 {
				res = append(res, []int{prev, nodes[j]})
			}
			prev = nodes[j]
			tmp--
		}
		for tmp > 0 {
			for vis[p] {
				p++
			}
			res = append(res, []int{prev, p})
			prev = p
			vis[p] = true
			tmp--
		}
	}

	res = append(res, []int{prev, n})

	return res
}

type Pair struct {
	first  int
	second int
}
