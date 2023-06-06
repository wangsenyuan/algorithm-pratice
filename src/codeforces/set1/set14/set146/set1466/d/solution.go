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

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		W := readNNums(reader, n)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, W, E)
		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d ", ans))
		}
		buf.WriteByte('\n')
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

func solve(n int, W []int, E [][]int) []int64 {
	// 同一个颜色的边要连在一起，组成一个connected component
	// 为了试k个颜色的value最大，应该让value越大的节点在边界处
	// 而且为了不出现单个节点的图，需要考虑节点的度数
	// 每次染色一条边，都是建该图分成了两部分
	deg := make([]int, n)
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		deg[u]++
		deg[v]++
	}

	vs := make([]*Vertex, n)

	res := make([]int64, n-1)

	for i := 0; i < n; i++ {
		res[0] += int64(W[i])
		vs[i] = &Vertex{W[i], deg[i]}
	}

	sort.Slice(vs, func(i, j int) bool {
		return vs[i].weight > vs[j].weight
	})
	var j int
	for i := 1; i < n-1; i++ {
		for j < n && vs[j].deg == 1 {
			j++
		}
		res[i] = res[i-1] + int64(vs[j].weight)
		vs[j].deg--
	}

	return res
}

type Vertex struct {
	weight int
	deg    int
}
