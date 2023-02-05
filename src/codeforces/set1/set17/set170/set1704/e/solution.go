package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		edges := make([][]int, m)
		for i := 0; i < m; i++ {
			edges[i] = readNNums(reader, 2)
		}
		res := solve(n, edges, A)

		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const M = 998244353

func solve(n int, edges [][]int, a []int) int {
	g := make([][]int, n+1)
	z := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		g[i] = make([]int, 0, 2)
		z[i] = make([]int, 0, 2)
	}

	D := make([]int, n+1)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g[u] = append(g[u], v)
		z[v] = append(z[v], u)
		D[u]++
	}
	A := make([]int, n+1)
	copy(A[1:], a)
	B := make([]int, n+1)
	copy(B, A)

	C := make([]int, n+1)

	for i := 1; i <= n; i++ {
		var flow bool
		for j := 1; j <= n; j++ {
			if B[j] > 0 {
				flow = true
				break
			}
		}
		if !flow {
			return i - 1
		}

		for u := 1; u <= n; u++ {
			if B[u] > 0 {
				// C[j] 可能被其他的j增加，所以不能直接等于
				C[u] += B[u] - 1
				for _, v := range g[u] {
					C[v]++
				}
			}
		}

		for u := 1; u <= n; u++ {
			if C[u] >= M {
				// 如果C[u] = M, B[u] = 0, 将得到错误的结果
				B[u] = C[u]%M + M
			} else {
				B[u] = C[u] % M
			}
			C[u] = 0
		}
	}

	// not ending before n iterations
	que := make([]int, n+1)
	var end int
	for u := 1; u <= n; u++ {
		if D[u] == 0 {
			end++
			que[end] = u
		}
	}

	s := make([]int, n+1)
	s[que[end]] = 1

	for i := 1; i <= end; i++ {
		s[que[i]] %= M
		for _, v := range z[que[i]] {
			D[v]--
			s[v] += s[que[i]]
			if D[v] == 0 {
				end++
				que[end] = v
			}
		}
	}

	var ans int64 = int64(n)

	for i := 1; i <= n; i++ {
		ans = (ans + int64(B[i])*int64(s[i])%M) % M
	}

	return int(ans)
}
