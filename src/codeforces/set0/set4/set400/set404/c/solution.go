package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)

	d := readNNums(reader, n)

	res := solve(n, k, d)

	if len(res) == 0 {
		fmt.Println("-1")
		return
	}

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
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

func solve(n int, k int, d []int) [][]int {
	deg := make([]int, n)

	at := make([][]int, n)

	for i := 0; i < n; i++ {
		at[i] = make([]int, 0, 1)
	}

	for i, x := range d {
		at[x] = append(at[x], i)
	}
	que := make([]int, n)
	var front, tail int
	var res [][]int

	addEdge := func(u int, v int) {
		res = append(res, []int{u + 1, v + 1})
		deg[u]++
		deg[v]++
	}
	for p := 0; p < n && front < n; p++ {
		if len(at[p]) == 0 {
			// 有一些节点没法挂上来
			return nil
		}
		if p == 0 {
			if len(at[0]) > 1 {
				// 有两个root了
				return nil
			}
			que[front] = at[0][0]
			front++
			continue
		}

		for _, u := range at[p] {
			for tail < front && (d[que[tail]] < p-1 || deg[que[tail]] == k) {
				tail++
			}
			if tail == front {
				return nil
			}
			addEdge(que[tail], u)
		}
		for _, u := range at[p] {
			que[front] = u
			front++
		}
	}

	return res
}
