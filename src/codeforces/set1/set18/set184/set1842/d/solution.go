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
	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 3)
	}
	time, sets, ts := solve(n, edges)
	if time < 0 {
		fmt.Println("inf")
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d %d\n", time, len(sets)))
	for i := 0; i < len(sets); i++ {
		buf.WriteString(fmt.Sprintf("%s %d\n", sets[i], ts[i]))
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

const inf = 1 << 60

func solve(n int, edges [][]int) (int, []string, []int) {
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = inf
		}
		dist[i][i] = 0
	}

	for _, e := range edges {
		u, v, w := e[0]-1, e[1]-1, e[2]
		dist[u][v] = w
		dist[v][u] = w
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	if dist[0][n-1] >= inf {
		return -1, nil, nil
	}

	var sets []string
	var times []int

	ord := make([]int, n)
	for i := 0; i < n; i++ {
		ord[i] = i
	}
	sort.Slice(ord, func(i, j int) bool {
		a, b := ord[i], ord[j]
		return dist[0][a] < dist[0][b] || dist[0][a] == dist[0][b] && a < b
	})

	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = '0'
	}
	for i := 0; i+1 < n; i++ {
		u, v := ord[i], ord[i+1]
		buf[u] = '1'
		sets = append(sets, string(buf))
		times = append(times, dist[0][v]-dist[0][u])
		if v == n-1 {
			break
		}
	}

	return dist[0][n-1], sets, times
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
