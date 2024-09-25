package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, v := readThreeNums(reader)
	res := solve(n, m, v)

	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer

	for _, e := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", e[0], e[1]))
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
func solve(n int, m int, v int) [][]int {
	n1 := n - 1
	if n1*(n1-1)/2+1 < m || m < n1 {
		return nil
	}

	// 这m条线都必须被使用掉
	// a * (a - 1) / 2 + b * (b - 1) / 2 == m
	nodes := make([]int, n+1)
	for i := 0; i <= n; i++ {
		nodes[i] = i
	}

	if v != 2 {
		nodes[2], nodes[v] = nodes[v], nodes[2]
	}

	res := make([][]int, 0, m)
	res = append(res, []int{nodes[1], nodes[2]})
	for i := 2; i < n && len(res) < m; i++ {
		for j := i + 1; j <= n && len(res) < m; j++ {
			res = append(res, []int{nodes[i], nodes[j]})
		}
	}
	return res
}
