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
		reader.ReadString('\n')
		N, L, K := readThreeNums(reader)
		S := make([]string, N)
		for i := 0; i < N; i++ {
			S[i] = readString(reader)
		}
		res, ops := solve(N, L, K, S)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		buf.WriteString(res)
		buf.WriteByte('\n')
		for _, op := range ops {
			buf.WriteString(fmt.Sprintf("%d %d\n", op[0], op[1]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

const INF = 1 << 30

func solve(N int, L int, K int, S []string) (string, [][]int) {

	// value[i][j] = concate(S[i], S[j]) get the common suf-pre length

	concate := func(u, v int) int {
		s, p := S[u], S[v]
		if s == p {
			p = p[:L-1]
		}
		ps := kmp(p)

		var j int
		for i := 0; i < L; {
			if s[i] == p[j] {
				i++
				j++
			}

			if i < L && (j == len(p) || s[i] != p[j]) {
				if j > 0 {
					j = ps[j-1]
				} else {
					i++
				}
			}
		}

		return L - j
	}

	dp_value := make([][]int, N)
	dp_path := make([][]int, N)

	for i := 0; i < N; i++ {
		dp_value[i] = make([]int, K)
		dp_path[i] = make([]int, K)
		dp_path[i][0] = i
	}

	cost := make([][]int, N)

	for i := 0; i < N; i++ {
		cost[i] = make([]int, N)
		for j := 0; j < N; j++ {
			cost[i][j] = concate(i, j)
		}
	}
	if K > 1 {
		for i := 0; i < N; i++ {
			prev_min := 0
			for j := 1; j < N; j++ {
				if cost[prev_min][i] > cost[j][i] {
					prev_min = j
				}
			}
			dp_value[i][1] = cost[prev_min][i]
			dp_path[i][1] = prev_min
		}
	}

	for i := 1; i+1 < K; i++ {
		for j := 0; j < N; j++ {
			cur := dp_value[j][i]
			for a := 0; a < N; a++ {
				next_value := cur + cost[j][a]
				if dp_value[a][i+1] == 0 || dp_value[a][i+1] > next_value {
					dp_value[a][i+1] = next_value
					dp_path[a][i+1] = j
				}
			}
		}
	}

	// pick k nodes, having the max value

	start := 0
	for i := 1; i < N; i++ {
		if dp_value[i][K-1] < dp_value[start][K-1] {
			start = i
		}
	}

	buf := make([]byte, dp_value[start][K-1]+L)
	res := make([][]int, K)

	for K > 0 {
		x := dp_value[start][K-1]
		res[K-1] = []int{start, x}
		copy(buf[x:x+L], []byte(S[start]))
		start = dp_path[start][K-1]
		K--
	}

	return string(buf), res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func kmp(s string) []int {
	n := len(s)
	lps := make([]int, n)

	for i := 1; i < n; i++ {
		j := lps[i-1]
		for j > 0 && s[j] != s[i] {
			j = lps[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		lps[i] = j
	}
	return lps
}
