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
	A, _ := reader.ReadString('\n')
	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = readNNums(reader, 2)
	}
	res := solve(n, k, A, C)
	fmt.Println(res)
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, k int, A string, C [][]int) string {

	P := make([]int, n)

	var dfs1 func(u int)

	var it int

	dfs1 = func(u int) {
		if C[u][0] > 0 {
			dfs1(C[u][0] - 1)
		}
		P[it] = u
		it++

		if C[u][1] > 0 {
			dfs1(C[u][1] - 1)
		}
	}
	dfs1(0)

	diff := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		diff[P[i]] = P[i]
		if i+1 < n {
			if A[P[i+1]] != A[P[i]] {
				diff[P[i]] = P[i+1]
			} else {
				diff[P[i]] = diff[P[i+1]]
			}
		}
	}

	dup := make([]bool, n)

	var dfs2 func(u, K int) int

	dfs2 = func(u, K int) int {
		if K == 0 {
			return 0
		}
		var res int

		if C[u][0] > 0 {
			res += dfs2(C[u][0]-1, K-1)
		}

		if res > 0 || A[u] < A[diff[u]] {
			// left need to duplicate, so do u
			dup[u] = true
			res++
		}
		if C[u][1] > 0 {
			// 如果u已经dup了, 右子树的处理和它就没有关系
			// 如果u不能dup
			if res > 0 {
				res += dfs2(C[u][1]-1, K-res)
			}
		}
		return res
	}

	dfs2(0, k)

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		u := P[i]
		buf.WriteByte(A[u])
		if dup[u] {
			buf.WriteByte(A[u])
		}
	}
	return buf.String()
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
