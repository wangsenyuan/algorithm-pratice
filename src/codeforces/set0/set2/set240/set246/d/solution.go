package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	c := readNNums(reader, n)
	edges := make([][]int, m)

	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 2)
	}

	res := solve(c, edges)

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

func solve(c []int, edges [][]int) int {
	ii := make(map[int]int)

	for _, x := range c {
		ii[x]++
	}

	arr := make([]int, 0, len(ii))
	for k := range ii {
		arr = append(arr, k)
	}

	n := len(arr)
	g := make([]map[int]bool, n)

	for i, x := range arr {
		ii[x] = i
		g[i] = make(map[int]bool)
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		u--
		v--
		if c[u] != c[v] {
			g[ii[c[u]]][ii[c[v]]] = true
			g[ii[c[v]]][ii[c[u]]] = true
		}
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		if len(g[i]) > len(g[res]) || len(g[i]) == len(g[res]) && arr[i] < arr[res] {
			res = i
		}
	}

	return arr[res]
}
