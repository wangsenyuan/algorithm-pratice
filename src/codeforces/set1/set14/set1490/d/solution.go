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
		A := readNNums(reader, n)
		res := solve(A)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func solve(A []int) []int {
	n := len(A)
	parent := make([]int, n)
	stack := make([]int, n)
	var p int
	for i := 0; i < n; i++ {
		for p > 0 && A[stack[p-1]] < A[i] {
			if p == 1 || A[stack[p-2]] > A[i] {
				parent[stack[p-1]] = i
			}
			p--
		}

		if p > 0 {
			parent[i] = stack[p-1]
		} else {
			parent[i] = -1
		}
		stack[p] = i
		p++
	}
	root := stack[0]
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		if parent[i] >= 0 {
			cnt[parent[i]]++
		}
	}
	out := make([][]int, n)
	for i := 0; i < n; i++ {
		out[i] = make([]int, 0, cnt[i])
	}

	for i := 0; i < n; i++ {
		if parent[i] >= 0 {
			out[parent[i]] = append(out[parent[i]], i)
		}
	}
	D := make([]int, n)
	var dfs func(u int)

	dfs = func(u int) {
		for _, v := range out[u] {
			D[v] = D[u] + 1
			dfs(v)
		}
	}
	dfs(root)

	return D
}
