package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	d, s := process(reader)

	res := solve(d, s)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
	}

	fmt.Print(buf.String())
}

func process(reader *bufio.Reader) ([]int, []int) {
	n := readNum(reader)
	d, s := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		d[i], s[i] = readTwoNums(reader)
	}
	return d, s
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

func solve(deg []int, s []int) [][]int {
	n := len(deg)

	que := make([]int, n)
	var head, tail int
	vis := make([]bool, n)

	for i := 0; i < n; i++ {
		if deg[i] == 1 {
			// 这个节点是叶子节点
			que[head] = i
			head++
		}
	}
	var res [][]int

	for tail < head {
		u := que[tail]
		tail++
		vis[u] = true
		if deg[u] == 0 {
			continue
		}
		v := s[u]
		if !vis[v] {
			res = append(res, []int{u, v})
		}
		deg[v]--
		// 要将u从v的贡献中去掉
		s[v] ^= u
		if deg[v] == 1 {
			que[head] = v
			head++
		}
	}

	return res
}
